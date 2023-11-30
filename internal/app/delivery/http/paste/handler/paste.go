package pastehandler

import (
	"errors"
	"html/template"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	pastedto "github.com/thansetan/kopas/internal/app/delivery/http/paste/dto"
	pasteusecase "github.com/thansetan/kopas/internal/app/usecase/paste"
	"github.com/thansetan/kopas/internal/helpers"
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

	c.Header("HX-Redirect", id)
	c.JSON(http.StatusCreated, gin.H{
		"id": id,
	})
}

func (h *pasteHandler) GetPasteByID(c *gin.Context) {
	id := c.Param("id")

	data, err := h.uc.GetPasteByID(c, id)
	if err != nil {
		errCode := http.StatusInternalServerError
		if errors.Is(err, helpers.ErrNotFound) {
			errCode = http.StatusNotFound
		}

		if strings.Contains(c.Request.Header.Get("Accept"), "text/html") && errCode == http.StatusNotFound {
			c.HTML(errCode, "not_found", nil)
			return
		}
		c.String(errCode, "%s", err.Error())

		return
	}

	htmlData := gin.H{
		"title":     data.Title,
		"content":   data.Content,
		"ExpiresIn": helpers.GetRemainingTime(data.ExpiresAt),
	}
	content, ok := helpers.HighlightCode(data.Content)

	if ok {
		htmlData["content"] = template.HTML(content)
	}

	switch c.Request.Header.Get("Accept") {
	case "application/json":
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	default:
		c.HTML(http.StatusOK, "view_paste", htmlData)
	}
}

func (h *pasteHandler) NewPaste(c *gin.Context) {
	c.HTML(http.StatusOK, "new_paste", gin.H{
		"content": "new_paste",
	})
}
