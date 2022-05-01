package services

import "resourcesAndFactories/pkg/domain/models"

func (s *Service) GetDashboardData() (models.Dashboard, error) {
	resources, err := s.getResources()
	if err != nil {
		return models.Dashboard{}, err
	}
	factories, err := s.getFactories()
	if err != nil {
		return models.Dashboard{}, err
	}
	return models.Dashboard{Resources: resources, Factories: factories}, nil
}
