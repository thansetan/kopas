package repository

import (
	"context"

	"github.com/thansetan/kopas/internal/domain/model"
)

type PasteRepository interface {
	Insert(context.Context, model.Paste) (string, error)
	GetByID(context.Context, string) (*model.Paste, error)
}
