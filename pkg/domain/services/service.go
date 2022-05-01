package services

import (
	"github.com/dgraph-io/badger/v3"
)

type Service struct {
	DB *badger.DB
}

func New(db *badger.DB) *Service {
	return &Service{
		DB: db,
	}
}
