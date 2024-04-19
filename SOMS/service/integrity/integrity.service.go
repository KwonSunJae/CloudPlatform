package integrity

import (
	"soms/repository"
	"soms/repository/integrity"
)

type IntegrityService struct {
	Repository *integrity.IntegrityRepository
}

var Service IntegrityService

func (s *IntegrityService) InitService() error {
	db, err := repository.OpenWithStorage()

	if err != nil {
		return err
	}

	s.Repository = &integrity.Repository
	s.Repository.AssignDB(db)

	return nil
}

func (s *IntegrityService) GetAllIntegrity() (*[]integrity.IntegrityRaw, error) {
	raws, err := s.Repository.GetAllIntegrity()

	return raws, err
}

func (s *IntegrityService) GetOneIntegrity(request_id string) (integrity.IntegrityRaw, error) {
	raw, err := s.Repository.GetOneIntegrity(request_id)

	return raw, err
}

func (s *IntegrityService) GetIntegrityByUserID(user_id string) (*[]integrity.IntegrityRaw, error) {
	raws, err := s.Repository.GetIntegrityByUserID(user_id)

	return raws, err
}

func (s *IntegrityService) CreateIntegrity(n integrity.IntegrityDto) (string, error) {
	id, DBSaveErr := s.Repository.InsertIntegrity(n)
	if DBSaveErr != nil {
		return "", DBSaveErr
	}
	return id, nil
}

func (s *IntegrityService) DeleteIntegrity(request_id string) (bool, error) {
	DBDeleteErr := s.Repository.DeleteIntegrity(request_id)
	if DBDeleteErr != nil {
		return false, DBDeleteErr
	}
	return true, nil
}

func (s *IntegrityService) UpdateIntegrity(n integrity.IntegrityDto) (bool, error) {
	DBUpdateErr := s.Repository.UpdateIntegrity(n)
	if DBUpdateErr != nil {
		return false, DBUpdateErr
	}
	return true, nil
}
