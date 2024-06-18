package deployment

import (
	"fmt"
	"os/exec"
	"soms/repository"
	"soms/repository/container/deployment"
	"soms/repository/user"
	resource "soms/util/resource/container/deployment"
)

type DeploymentService struct {
	Repository *deployment.DeploymentRepository
}

var Service DeploymentService

func (s *DeploymentService) InitService() error {
	db, err := repository.OpenWithFile()

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

func (s *DeploymentService) EnrollDeployment(n deployment.DeploymentDto) error {

	_, err3 := s.Repository.InsertDeployment(n)
	if err3 != nil {
		return fmt.Errorf("db error: %v", err3)
	}

	return nil
}

func (s *DeploymentService) ApproveDeployment(id string) error {

	n, err := s.Repository.GetOneDeployment(id)
	if err != nil {
		return fmt.Errorf("해당 데이터가 없음: %v", err)
	}

	targetUser, err := user.Repository.GetOneUserByUUID(n.UUID)
	if err != nil {
		return fmt.Errorf("해당 데이터가 없음: %v", err)
	}

	deploymentManager := resource.New()
	err2 := deploymentManager.UserID(targetUser.UserID).ApiVersion(n.ApiVersion).Kind(n.Kind).MetadataName(n.MetadataName).MetadataLabelsApp(n.MetadataLabelsApp).SpecReplicas(n.SpecReplicas).SpecSelectorMatchlabelsApp(n.SpecSelectorMatchlabelsApp).SpecTemplateMetadataLabelsApp(n.SpecTemplateMetadataLabelsApp).SpecTemplateSpecContainersName(n.SpecTemplateSpecContainersName).SpecTemplateSpecContainersImage(n.SpecTemplateSpecContainersImage).SpecTemplateSpecContainersPortsContainerport(n.SpecTemplateSpecContainersPortsContainerport).Build()
	if err2 != nil {
		return err2
	}

	var approvedDeployment deployment.DeploymentDto
	approvedDeployment.Status = "Approved"

	_, err4 := s.Repository.UpdateOneDeployment(id, approvedDeployment)
	if err4 != nil {
		fmt.Println("db error: %v", err4)
		return fmt.Errorf("db error: %v", err4)
	}

	return nil
}

func (s *DeploymentService) UpdateDeployment(id string, n deployment.DeploymentDto) error {
	dpData, err0 := s.Repository.GetOneDeployment(id)
	if err0 != nil {
		return fmt.Errorf("해당 데이터가 없음: %v", err0)
	}
	targetUser, err := user.Repository.GetOneUserByUUID(dpData.UUID)
	if err != nil {
		return fmt.Errorf("해당 데이터가 없음: %v", err)
	}

	cmd := exec.Command("kubectl", "delete", "deployment", dpData.MetadataName, "-n", targetUser.UserID) // , "-n", dpData.UserID
	_, err2 := cmd.CombinedOutput()
	if err2 != nil {
		return fmt.Errorf("기존 deployment 삭제실패: %v", err2)
	}

	deploymentManager := resource.New()
	err3 := deploymentManager.UserID(targetUser.UserID).ApiVersion(n.ApiVersion).Kind(n.Kind).MetadataName(n.MetadataName).MetadataLabelsApp(n.MetadataLabelsApp).SpecReplicas(n.SpecReplicas).SpecSelectorMatchlabelsApp(n.SpecSelectorMatchlabelsApp).SpecTemplateMetadataLabelsApp(n.SpecTemplateMetadataLabelsApp).SpecTemplateSpecContainersName(n.SpecTemplateSpecContainersName).SpecTemplateSpecContainersImage(n.SpecTemplateSpecContainersImage).SpecTemplateSpecContainersPortsContainerport(n.SpecTemplateSpecContainersPortsContainerport).Build()
	if err3 != nil {
		return err3
	}

	_, err4 := s.Repository.UpdateOneDeployment(id, n)
	if err4 != nil {
		return fmt.Errorf("db error: %v", err)
	}
	return nil
}

func (s *DeploymentService) DeleteDeployment(id string) error {
	dpData, err0 := s.Repository.GetOneDeployment(id)

	if err0 != nil {
		return fmt.Errorf("해당 데이터가 없음: %v", err0)
	}

	targetUser, err := user.Repository.GetOneUserByUUID(dpData.UUID)
	if err != nil {
		return fmt.Errorf("해당 데이터가 없음: %v", err)
	}

	cmd := exec.Command("kubectl", "delete", "deployment", dpData.MetadataName, "-n", targetUser.UserID) // , "-n", dpData.UserID
	output, err2 := cmd.CombinedOutput()
	if err2 != nil {
		return fmt.Errorf("삭제실패: %v", err2)
	}
	_, err3 := s.Repository.DeleteOneDeployment(id)
	if err3 != nil {
		fmt.Print(output)
		return fmt.Errorf("kubectl 명령 실행 중 오류 발생: %v", err2)

	}
	return err
}

func (s *DeploymentService) GetDeploymentsStatus(uuid string) (string, error) {
	targetUser, err := user.Repository.GetOneUserByUUID(uuid)
	if err != nil {
		return "", fmt.Errorf("해당 데이터가 없음: %v", err)
	}
	// kubectl 명령 실행

	cmd := exec.Command("kubectl", "get", "deployments", "-o", "json", "-n", targetUser.UserID)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("kubectl 명령 실행 중 오류 발생: %v", err)
	}

	return string(output), nil
}
