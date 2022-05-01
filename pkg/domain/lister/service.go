package lister

import (
	"bytes"
	"encoding/gob"
	"github.com/dgraph-io/badger/v3"
	"resourcesAndFactories/pkg/domain/models"
)

type Service interface {
	getResources() (map[models.ResourceType]models.Resource, error)
}

type service struct {
	*badger.DB
}

func New(db *badger.DB) Service {
	return &service{db}
}

func (s *service) getResources() (resources map[models.ResourceType]models.Resource, err error) {
	err = s.DB.View(func(txn *badger.Txn) error {
		for _, resourceType := range []models.ResourceType{models.IronResource, models.CopperResource, models.GoldResource} {
			serializedResource, err := txn.Get([]byte(resourceType))
			if err != nil {
				return err
			}
			val, err := serializedResource.ValueCopy(nil)
			if err != nil {
				return err
			}
			d := gob.NewDecoder(bytes.NewReader(val))
			var resource models.Resource
			if err := d.Decode(&resource); err != nil {
				return err
			}
			resources[resourceType] = resource
		}
		return nil
	})
	return
}
