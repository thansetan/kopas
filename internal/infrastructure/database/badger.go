package database

import "github.com/dgraph-io/badger/v4"

func NewBadger(dbPath string) (*badger.DB, error) {
	db, err := badger.Open(badger.DefaultOptions(dbPath))
	if err != nil {
		return nil, err
	}

	return db, nil
}
