package deployment

import (
"soms/repository"
"soms/repository/deployment"
)

type DeploymentService struct {
	Repository * deployment.DeploymentRepository
}

var Service DeploymentService

func (s *DeploymentService) InitService() error {
	db, err := repository.OpenWithMemory()

	if err != nil {
		return err
	}

	s.Repository = &deployment.Repository
	s.Repository.AssignDB(db)

	return nil
}

func (s *DeploymentService) GetAllDeployment() (*[]deployment.DeploymentRaw, error) {
	raws, err := s.Repository.GetAllDeployment()

	return raws, err
}

func (s *DeploymentService) GetOneDeployment(id string) (*deployment.DeploymentRaw, error) {
	raw, err := s.Repository.GetOneDeployment(id)

	return raw, err
}

func (s *DeploymentService) CreateDeployment(n deployment.DeploymentDto) error {
	_, err := s.Repository.InsertDeployment(n)

	return err
}

func (s *DeploymentService) UpdateDeployment(id string, n deployment.DeploymentDto) error {
	_, err := s.Repository.UpdateOneDeployment(id, n)

	return err
}

func (s *DeploymentService) DeleteDeployment(id string) error {
	_, err := s.Repository.DeleteOneDeployment(id)

	return err
}
