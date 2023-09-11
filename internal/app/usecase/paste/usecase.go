package pasteusecase

import (
	"errors"

	pastedto "github.com/thansetan/kopas/internal/app/delivery/http/paste/dto"
	"github.com/thansetan/kopas/internal/domain/model"
	"github.com/thansetan/kopas/internal/domain/repository"
	"github.com/thansetan/kopas/internal/helpers"
)

type PasteUsecase interface {
	NewPaste(data pastedto.Paste) (string, error)
	GetPasteByID(id string) (*pastedto.Paste, error)
}

type pasteUsecase struct {
	repo repository.PasteRepository
}

func NewPasteUsecase(repo repository.PasteRepository) PasteUsecase {
	return &pasteUsecase{
		repo: repo,
	}
}

func (uc *pasteUsecase) NewPaste(data pastedto.Paste) (string, error) {
	pasteData := model.Paste{
		Title:   []byte(data.Title),
		Content: []byte(data.Content),
	}

	if !helpers.ValidSize(pasteData.Content) {
		return "", errors.New("size can't be more than 20MB")
	}

	id, err := uc.repo.Insert(pasteData)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (uc *pasteUsecase) GetPasteByID(id string) (*pastedto.Paste, error) {
	data, err := uc.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	pasteData := pastedto.Paste{
		Title:   string(data.Title),
		Content: string(data.Content),
	}

	return &pasteData, nil
}
