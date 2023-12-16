package service

import (
	"soms/repository"
	"soms/repository/container/service"
)

type ServiceService struct {
	Repository *service.ServiceRepository
}

var Service ServiceService

func (s *ServiceService) InitService() error {
	db, err := repository.OpenWithMemory()

	if err != nil {
		return err
	}

	s.Repository = &service.Repository
	s.Repository.AssignDB(db)

	return nil
}

func (s *ServiceService) GetAllService() (*[]service.ServiceRaw, error) {
	raws, err := s.Repository.GetAllService()

	return raws, err
}

func (s *ServiceService) GetOneService(id string) (*service.ServiceRaw, error) {
	raw, err := s.Repository.GetOneService(id)

	return raw, err
}

func (s *ServiceService) CreateService(n service.ServiceDto) error {
	_, err := s.Repository.InsertService(n)

	return err
}

func (s *ServiceService) UpdateService(id string, n service.ServiceDto) error {
	_, err := s.Repository.UpdateOneService(id, n)

	return err
}

func (s *ServiceService) DeleteService(id string) error {
	_, err := s.Repository.DeleteOneService(id)

	return err
}
