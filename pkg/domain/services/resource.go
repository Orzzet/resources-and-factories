package services

import (
	"bytes"
	"encoding/gob"
	"github.com/dgraph-io/badger/v3"
	"resourcesAndFactories/pkg/domain/models"
)

func (s *Service) getResources() (resources []models.Resource, err error) {
	for _, resourceType := range []models.ResourceType{models.IronResource, models.CopperResource, models.GoldResource} {
		resource, err := s.getResource(resourceType)
		if err != nil {
			return []models.Resource{}, err
		}
		resources = append(resources, resource)
	}
	return
}

func (s *Service) initializeResources() (err error) {
	for _, resourceType := range []models.ResourceType{models.IronResource, models.CopperResource, models.GoldResource} {
		_, err = s.getResource(resourceType)
		if err == nil {
			continue
		}
		err = s.setResource(resourceType, 0)
	}
	return
}

func (s *Service) getResource(resourceType models.ResourceType) (resource models.Resource, err error) {
	err = s.DB.View(func(txn *badger.Txn) error {
		serializedResource, err := txn.Get([]byte(resourceType))
		if err != nil {
			return err
		}
		val, err := serializedResource.ValueCopy(nil)
		if err != nil {
			return err
		}
		d := gob.NewDecoder(bytes.NewReader(val))
		if err := d.Decode(&resource); err != nil {
			return err
		}
		return nil
	})
	return
}

func (s *Service) setResource(resourceType models.ResourceType, amount float64) (err error) {
	return s.DB.Update(func(txn *badger.Txn) error {
		resource := models.Resource{
			ResourceType: resourceType,
			Amount:       amount,
		}
		var b bytes.Buffer
		e := gob.NewEncoder(&b)
		if err := e.Encode(resource); err != nil {
			return err
		}
		err = s.DB.Update(func(txn *badger.Txn) error {
			err := txn.Set([]byte(resource.ResourceType), b.Bytes())
			return err
		})
		if err != nil {
			return err
		}
		return nil
	})
}
