package pastehandler

import (
	"errors"
	"net/http"

	"github.com/dgraph-io/badger/v4"
	"github.com/gin-gonic/gin"
	pastedto "github.com/thansetan/kopas/internal/app/delivery/http/paste/dto"
	pasteusecase "github.com/thansetan/kopas/internal/app/usecase/paste"
)

type PasteHandler interface {
	InsertPaste(c *gin.Context)
	GetPasteByID(c *gin.Context)
	NewPaste(c *gin.Context)
}

type pasteHandler struct {
	uc pasteusecase.PasteUsecase
}

func NewPasteHandler(uc pasteusecase.PasteUsecase) PasteHandler {
	return &pasteHandler{
		uc: uc,
	}
}

func (h *pasteHandler) InsertPaste(c *gin.Context) {
	var pasteData pastedto.PasteReq
	err := c.ShouldBind(&pasteData)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	id, err := h.uc.NewPaste(c, pasteData)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": id,
	})
}

func (h *pasteHandler) GetPasteByID(c *gin.Context) {
	id := c.Param("id")

	data, err := h.uc.GetPasteByID(c, id)
	if err != nil {
		errCode := http.StatusInternalServerError
		if errors.Is(err, badger.ErrKeyNotFound) {
			errCode = http.StatusNotFound
		}
		c.String(errCode, "%s", err.Error())
		return
	}

	switch c.Request.Header.Get("Accept") {
	case "application/json":
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	default:
		c.HTML(http.StatusOK, "paste.html", gin.H{
			"title":     data.Title,
			"content":   data.Content,
			"expiresAt": data.ExpiresAt,
		})
	}
}

func (h *pasteHandler) NewPaste(c *gin.Context) {
	c.HTML(http.StatusOK, "new.html", nil)
}
