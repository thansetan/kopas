package httproute

import (
	"github.com/dgraph-io/badger/v4"
	"github.com/gin-gonic/gin"
	"github.com/thansetan/kopas/internal/app/delivery/http/paste"
)

func NewRoute(db *badger.DB) *gin.Engine {
	r := gin.Default()

	paste.Route(r, db)

	return r
}
