package badger

import (
	"github.com/dgraph-io/badger/v3"
)

func New() (*badger.DB, error) {
	db, err := badger.Open(badger.DefaultOptions("/tmp/badger"))
	if err != nil {
		return nil, err
	}
	return db, err
}
