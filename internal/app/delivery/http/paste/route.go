package paste

import (
	"fmt"
	"time"

	"github.com/dgraph-io/badger/v4"
	"github.com/gin-gonic/gin"
	"github.com/thansetan/kopas/internal/app/delivery/http/middlewares"
	pastehandler "github.com/thansetan/kopas/internal/app/delivery/http/paste/handler"
	pasterepository "github.com/thansetan/kopas/internal/app/repository/paste"
	pasteusecase "github.com/thansetan/kopas/internal/app/usecase/paste"
	"github.com/thansetan/kopas/pkg/helpers"
)

func Route(r *gin.Engine, db *badger.DB) {
	repo := pasterepository.NewPasteRepository(db)
	uc := pasteusecase.NewPasteUsecase(repo)
	handler := pastehandler.NewPasteHandler(uc)

	rl := middlewares.NewLimiter(10, 24*time.Hour) // create 10 paste every day

	dir, err := helpers.GetCurrentFileDir()
	if err != nil {
		panic(err)
	}

	r.LoadHTMLGlob(fmt.Sprintf("%s/views/*", dir))

	r.GET("", handler.NewPaste)
	r.GET("/:id", handler.GetPasteByID)
	r.Use(rl.RateLimitMiddleware()).POST("/paste", handler.InsertPaste)
}
