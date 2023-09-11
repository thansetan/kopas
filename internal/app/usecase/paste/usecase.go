package pasteusecase

import (
	"errors"
	"time"

	pastedto "github.com/thansetan/kopas/internal/app/delivery/http/paste/dto"
	"github.com/thansetan/kopas/internal/domain/model"
	"github.com/thansetan/kopas/internal/domain/repository"
	"github.com/thansetan/kopas/internal/helpers"
)

type PasteUsecase interface {
	NewPaste(data pastedto.PasteReq) (string, error)
	GetPasteByID(id string) (*pastedto.PasteResp, error)
}

type pasteUsecase struct {
	repo repository.PasteRepository
}

func NewPasteUsecase(repo repository.PasteRepository) PasteUsecase {
	return &pasteUsecase{
		repo: repo,
	}
}

func (uc *pasteUsecase) NewPaste(data pastedto.PasteReq) (string, error) {
	pasteData := model.Paste{
		Title:   []byte(data.Title),
		Content: []byte(data.Content),
	}

	expDur := helpers.GetTTL(data.ExpiresAt)
	switch expDur {
	case 0:
		pasteData.ExpiresAt = 0
	default:
		pasteData.ExpiresAt = time.Now().Add(expDur).Unix()
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

func (uc *pasteUsecase) GetPasteByID(id string) (*pastedto.PasteResp, error) {
	data, err := uc.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	pasteData := pastedto.PasteResp{
		Title:     string(data.Title),
		Content:   string(data.Content),
		ExpiresAt: data.ExpiresAt,
	}

	return &pasteData, nil
}
