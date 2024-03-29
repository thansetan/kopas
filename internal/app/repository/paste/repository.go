package pasterepository

import (
	"bytes"
	"context"
	"encoding/gob"
	"errors"

	"github.com/dgraph-io/badger/v4"
	"github.com/thansetan/kopas/internal/domain/model"
	"github.com/thansetan/kopas/internal/domain/repository"
	"github.com/thansetan/kopas/internal/helpers"
)

type pasteRepository struct {
	db *badger.DB
}

func NewPasteRepository(db *badger.DB) repository.PasteRepository {
	return &pasteRepository{
		db: db,
	}
}

func (repo *pasteRepository) Insert(ctx context.Context, data model.Paste) (string, error) {
	txn := repo.db.NewTransaction(true)
	defer txn.Discard()

	id, err := helpers.GenerateID()
	if err != nil {
		return "", err
	}

	var generateIDCount int
	_, exists := repo.GetByID(ctx, id) // check if paste with generated ID already exists
	for exists == nil {                // no error means the paste exists
		if generateIDCount > 4 {
			return "", errors.New("failed to generate ID")
		}
		id, err = helpers.GenerateID() // generate newID
		if err != nil {
			return "", err
		}
		_, exists = repo.GetByID(ctx, id) // check again
		generateIDCount++
	}

	var b bytes.Buffer
	err = gob.NewEncoder(&b).Encode(data)
	if err != nil {
		return "", err
	}

	e := &badger.Entry{
		Key:       []byte(id),
		Value:     b.Bytes(),
		ExpiresAt: uint64(data.ExpiresAt),
	}

	err = txn.SetEntry(e)
	if err != nil {
		return "", err
	}

	if err := txn.Commit(); err != nil {
		return "", err
	}

	return id, nil
}

func (repo *pasteRepository) GetByID(ctx context.Context, id string) (*model.Paste, error) {
	pasteData := new(model.Paste)

	err := repo.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(id))
		if err != nil {
			return err
		}

		err = item.Value(func(val []byte) error {
			var b bytes.Buffer
			b.WriteString(string(val))
			d := gob.NewDecoder(&b)
			err := d.Decode(pasteData)
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return pasteData, nil
}
