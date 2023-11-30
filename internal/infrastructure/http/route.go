package httproute

import (
	"strings"

	"github.com/dgraph-io/badger/v4"
	"github.com/gin-gonic/gin"
	"github.com/thansetan/kopas/internal/app/delivery/http/paste"
)

func NewRoute(db *badger.DB) *gin.Engine {
	r := gin.Default()

	paste.Route(r, db)

	r.NoRoute(func(ctx *gin.Context) {
		if strings.Contains(ctx.Request.Header.Get("Accept"), "text/html") {
			ctx.HTML(404, "not_found", nil)
		}
	})
	return r
}
