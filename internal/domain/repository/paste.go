package repository

import "github.com/thansetan/kopas/internal/domain/model"

type PasteRepository interface {
	Insert(model.Paste) (string, error)
	GetByID(string) (*model.Paste, error)
}
