package pasteusecase

import (
	"context"
	"errors"
	"time"

	"github.com/dgraph-io/badger/v4"
	pastedto "github.com/thansetan/kopas/internal/app/delivery/http/paste/dto"
	"github.com/thansetan/kopas/internal/domain/model"
	"github.com/thansetan/kopas/internal/domain/repository"
	"github.com/thansetan/kopas/internal/helpers"
)

type PasteUsecase interface {
	NewPaste(context.Context, pastedto.PasteReq) (string, error)
	GetPasteByID(context.Context, string) (*pastedto.PasteResp, error)
}

type pasteUsecase struct {
	repo repository.PasteRepository
}

func NewPasteUsecase(repo repository.PasteRepository) PasteUsecase {
	return &pasteUsecase{
		repo: repo,
	}
}

func (uc *pasteUsecase) NewPaste(ctx context.Context, data pastedto.PasteReq) (string, error) {
	compressedContent, err := helpers.Compress([]byte(data.Content))
	if err != nil {
		return "", err
	}
	pasteData := model.Paste{
		Title:   []byte(data.Title),
		Content: compressedContent,
	}

	expDur := helpers.GetTTL(data.ExpiresAt)
	switch expDur {
	case 0:
		pasteData.ExpiresAt = 0
	default:
		pasteData.ExpiresAt = time.Now().Add(expDur).Unix()
	}

	if !helpers.IsValidSize(pasteData.Content) {
		return "", errors.New("size can't be more than 20MB")
	}

	id, err := uc.repo.Insert(ctx, pasteData)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (uc *pasteUsecase) GetPasteByID(ctx context.Context, id string) (*pastedto.PasteResp, error) {
	data, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, badger.ErrKeyNotFound) {
			return nil, helpers.ErrNotFound
		}
		return nil, err
	}

	decompressedContent, err := helpers.Decompress(data.Content)
	if err != nil {
		return nil, err
	}
	pasteData := pastedto.PasteResp{
		Title:     string(data.Title),
		Content:   string(decompressedContent),
		ExpiresAt: data.ExpiresAt,
	}

	return &pasteData, nil
}
