package badger

import (
	"github.com/dgraph-io/badger/v3"
)

func New() (*badger.DB, error) {
	db, err := badger.Open(badger.DefaultOptions("/tmp/badger"))
	if err != nil {
		return nil, err
	}
	updates := make(map[string]string)
	txn := db.NewTransaction(true)
	for k, v := range updates {
		if err := txn.Set([]byte(k), []byte(v)); err == badger.ErrTxnTooBig {
			_ = txn.Commit()
			txn = db.NewTransaction(true)
			_ = txn.Set([]byte(k), []byte(v))
		}
	}
	_ = txn.Commit()
	return db, err
}
