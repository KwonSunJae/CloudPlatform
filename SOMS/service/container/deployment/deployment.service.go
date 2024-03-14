package deployment

import (
	"fmt"
	"os/exec"
	"soms/repository"
	"soms/repository/container/deployment"
	resource "soms/util/resource/container/deployment"
)

type DeploymentService struct {
	Repository *deployment.DeploymentRepository
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
	deploymentManager := resource.New()
	err := deploymentManager.UserID(n.UserID).ApiVersion(n.ApiVersion).Kind(n.Kind).MetadataName(n.MetadataName).MetadataLabelsApp(n.MetadataLabelsApp).SpecReplicas(n.SpecReplicas).SpecSelectorMatchlabelsApp(n.SpecSelectorMatchlabelsApp).SpecTemplateMetadataLabelsApp(n.SpecTemplateMetadataLabelsApp).SpecTemplateSpecContainersName(n.SpecTemplateSpecContainersName).SpecTemplateSpecContainersImage(n.SpecTemplateSpecContainersImage).SpecTemplateSpecContainersPortsContainerport(n.SpecTemplateSpecContainersPortsContainerport).Build()
	if err != nil {
		return err
	}
	_, err2 := s.Repository.InsertDeployment(n)
	if err2 != nil {
		return fmt.Errorf("db error: %v", err2)
	}

	return nil
}

func (s *DeploymentService) UpdateDeployment(id string, n deployment.DeploymentDto) error {
	dpData, err0 := s.Repository.GetOneDeployment(id)
	if err0 != nil {
		return fmt.Errorf("해당 데이터가 없음: %v", err0)
	}

	cmd := exec.Command("kubectl", "delete", "deployment", dpData.MetadataName, "-n", dpData.UserID) // , "-n", dpData.UserID
	_, err2 := cmd.CombinedOutput()
	if err2 != nil {
		return fmt.Errorf("기존 deployment 삭제실패: %v", err2)
	}

	deploymentManager := resource.New()
	err3 := deploymentManager.UserID(n.UserID).ApiVersion(n.ApiVersion).Kind(n.Kind).MetadataName(n.MetadataName).MetadataLabelsApp(n.MetadataLabelsApp).SpecReplicas(n.SpecReplicas).SpecSelectorMatchlabelsApp(n.SpecSelectorMatchlabelsApp).SpecTemplateMetadataLabelsApp(n.SpecTemplateMetadataLabelsApp).SpecTemplateSpecContainersName(n.SpecTemplateSpecContainersName).SpecTemplateSpecContainersImage(n.SpecTemplateSpecContainersImage).SpecTemplateSpecContainersPortsContainerport(n.SpecTemplateSpecContainersPortsContainerport).Build()
	if err3 != nil {
		return err3
	}

	_, err := s.Repository.UpdateOneDeployment(id, n)
	if err != nil {
		return fmt.Errorf("db error: %v", err)
	}
	return nil
}

func (s *DeploymentService) DeleteDeployment(id string) error {
	dpData, err0 := s.Repository.GetOneDeployment(id)
	if err0 != nil {
		return fmt.Errorf("해당 데이터가 없음: %v", err0)
	}

	cmd := exec.Command("kubectl", "delete", "deployment", dpData.MetadataName, "-n", dpData.UserID) // , "-n", dpData.UserID
	output, err2 := cmd.CombinedOutput()
	if err2 != nil {
		return fmt.Errorf("삭제실패: %v", err2)
	}
	_, err := s.Repository.DeleteOneDeployment(id)
	if err != nil {
		fmt.Print(output)
		return fmt.Errorf("kubectl 명령 실행 중 오류 발생: %v", err2)

	}
	return err
}

func (s *DeploymentService) GetDeploymentsStatus() (string, error) {
	// kubectl 명령 실행
	cmd := exec.Command("kubectl", "get", "deployments", "-o", "json", "-n", "test")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("kubectl 명령 실행 중 오류 발생: %v", err)
	}

	return string(output), nil
}
