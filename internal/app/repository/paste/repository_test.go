package pasterepository

import (
	"context"
	"testing"

	"github.com/dgraph-io/badger/v4"
	"github.com/thansetan/kopas/internal/domain/model"
	"github.com/thansetan/kopas/pkg/helpers"
)

var repo *pasteRepository

func TestMain(m *testing.M) {
	badgerPath := helpers.GetEnvOrDefault("BADGER_PATH", "badger")
	db, _ := badger.Open(badger.DefaultOptions(badgerPath))
	defer db.Close()
	repo = &pasteRepository{
		db: db,
	}
	m.Run()
}
func TestPaste(t *testing.T) {
	var (
		content = `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nam condimentum egestas auctor. 
Nunc commodo, magna nec sollicitudin venenatis, ex nisi bibendum nibh, ut malesuada ipsum est vel enim.
Quisque vel lectus non elit ultricies bibendum. Fusce vehicula volutpat est nec bibendum. Praesent ipsum sapien,
eleifend sit amet malesuada a, efficitur sit amet ex.
Proin vitae fringilla arcu. Aliquam lobortis blandit urna, in aliquet quam tempus sed.
Sed suscipit vel ante non luctus. Etiam et mattis ex.`
		id string
	)

	data := model.Paste{
		Content: []byte(content),
	}

	t.Run("add new paste", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		pasteID, err := repo.Insert(ctx, data)
		if err != nil {
			t.Error("failed to insert data, err: ", err)
		}
		id = pasteID
	})

	t.Run("get inserted paste", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		pasteData := new(model.Paste)
		pasteData, err := repo.GetByID(ctx, id)
		if err != nil {
			t.Error("failed to get paste data, err: ", err)
		}
		if string(pasteData.Content) != content {
			t.Errorf("get paste data success, but the data is different.\nexpected: %s\ngot: %s\n", content, string(pasteData.Content))
		}
	})

	t.Run("get non-existent paste", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		pasteData := new(model.Paste)
		pasteData, err := repo.GetByID(ctx, "")
		if err == nil {
			t.Errorf("test failed, expected: nothing, but got: %s\n", pasteData.Content)
		}
	})
}
