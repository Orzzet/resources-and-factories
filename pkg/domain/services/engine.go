package services

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
	factories, err := s.getFactories()
	for _, factory := range factories {
		resource, err := s.getResource(factory.Production.ResourceType)
		if err != nil {
			return err
		}
		err = s.setResource(
			resource.ResourceType,
			resource.Amount+factory.FactoryStats.Production.Amount,
		)
		if err != nil {
			return err
		}
	}
	return
}
