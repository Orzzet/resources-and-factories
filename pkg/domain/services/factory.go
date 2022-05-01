package services

import (
	"bytes"
	"encoding/gob"
	"github.com/dgraph-io/badger/v3"
	"resourcesAndFactories/pkg/domain/models"
)

func (s *Service) getFactories() (factories []models.Factory, err error) {
	factoriesStats := models.GetFactoriesStats()
	err = s.DB.View(func(txn *badger.Txn) error {
		for _, factoryType := range []models.FactoryType{models.IronFactory, models.CopperFactory, models.GoldFactory} {
			serializedFactory, err := txn.Get([]byte(factoryType))
			if err != nil {
				return err
			}
			val, err := serializedFactory.ValueCopy(nil)
			if err != nil {
				return err
			}
			d := gob.NewDecoder(bytes.NewReader(val))
			var factory models.Factory
			if err := d.Decode(&factory); err != nil {
				return err
			}
			factory.FactoryStats = factoriesStats[factory.FactoryType][int(factory.Level)]
			factories = append(factories, factory)
		}
		return nil
	})
	return
}

func (s *Service) initializeFactories() (err error) {
	return s.DB.Update(func(txn *badger.Txn) error {
		for _, factoryType := range []models.FactoryType{models.IronFactory, models.CopperFactory, models.GoldFactory} {
			_, err := txn.Get([]byte(factoryType))
			if err == nil {
				continue
			}
			factory := models.Factory{
				FactoryType: factoryType,
				Level:       1,
			}
			var b bytes.Buffer
			e := gob.NewEncoder(&b)
			if err := e.Encode(factory); err != nil {
				return err
			}
			err = s.DB.Update(func(txn *badger.Txn) error {
				err := txn.Set([]byte(factoryType), b.Bytes())
				return err
			})
			if err != nil {
				return err
			}
		}
		return err
	})
}
