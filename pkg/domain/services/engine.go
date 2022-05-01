package services

import (
	"fmt"
	"resourcesAndFactories/pkg/domain/models"
)

func (s *Service) EngineStartup() (err error) {
	err = s.initializeResources()
	if err != nil {
		return err
	}
	err = s.initializeFactories()
	if err != nil {
		return err
	}
	return
}

func (s *Service) EngineTick() (err error) {
	fmt.Println("tick")
	return
	factories, err := s.getFactories()
	factoriesStats := models.GetFactoriesStats()
	for _, factory := range factories {
		resource, err := s.getResource(factory.Production.ResourceType)
		if err != nil {
			return err
		}
		err = s.setResource(
			resource.ResourceType,
			resource.Amount+factoriesStats[factory.FactoryType][int(factory.Level)].Production.Amount,
		)
		if err != nil {
			return err
		}
	}
	return
}
