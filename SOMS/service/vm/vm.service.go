package vm

import (
	"soms/repository"
	"soms/repository/vm"
)

type VmService struct {
	Repository *vm.VmRepository
}

var Service VmService

func (s *VmService) InitService() error {
	db, err := repository.OpenWithMemory()

	if err != nil {
		return err
	}

	s.Repository = &vm.Repository
	s.Repository.AssignDB(db)

	return nil
}

func (s *VmService) GetAllVm() (*[]vm.VmRaw, error) {
	raws, err := s.Repository.GetAllVm()

	return raws, err
}

func (s *VmService) GetOneVm(id string) (*vm.VmRaw, error) {
	raw, err := s.Repository.GetOneVm(id)

	return raw, err
}

func (s *VmService) CreateVm(n vm.VmDto) error {
	_, err := s.Repository.InsertVm(n)

	return err
}

func (s *VmService) UpdateVm(id string, n vm.VmDto) error {
	_, err := s.Repository.UpdateOneVm(id, n)

	return err
}

func (s *VmService) DeleteVm(id string) error {
	_, err := s.Repository.DeleteOneVm(id)

	return err
}
