package replicaset

import (
	"fmt"
	"os/exec"
	"soms/repository"
	"soms/repository/container/replicaset"
	resource "soms/util/resource/container/replicaset"
)

type ReplicasetService struct {
	Repository *replicaset.ReplicasetRepository
}

var Service ReplicasetService

func (s *ReplicasetService) InitService() error {
	db, err := repository.OpenWithMemory()

	if err != nil {
		return err
	}

	s.Repository = &replicaset.Repository
	s.Repository.AssignDB(db)

	return nil
}

func (s *ReplicasetService) GetAllReplicaset() (*[]replicaset.ReplicasetRaw, error) {
	raws, err := s.Repository.GetAllReplicaset()

	return raws, err
}

func (s *ReplicasetService) GetOneReplicaset(id string) (*replicaset.ReplicasetRaw, error) {
	raw, err := s.Repository.GetOneReplicaset(id)

	return raw, err
}

func (s *ReplicasetService) CreateReplicaset(n replicaset.ReplicasetDto) error {
	replicasetManager := resource.New()
	err := replicasetManager.UserId("test").ApiVersion(n.ApiVersion).Kind(n.Kind).MetadataName(n.MetadataName).SpecReplicas(n.SpecReplicas).SpecSelectorMatchlabelsApp(n.SpecSelectorMatchlabelsApp).SpecTemplateMetadataName(n.SpecTemplateMetadataName).SpecTemplateMetadataLabelsApp(n.SpecTemplateMetadataLabelsApp).SpecTemplateSpecContainersName(n.SpecTemplateSpecContainersName).SpecTemplateSpecContainersImage(n.SpecTemplateSpecContainersImage).SpecTemplateSpecContainersPortsContainerport(n.SpecTemplateSpecContainersPortsContainerport).Build()
	if err != nil {
		return err
	}
	_, err2 := s.Repository.InsertReplicaset(n)
	if err2 != nil {
		return fmt.Errorf("db error: %v", err2)
	}

	return nil
}

func (s *ReplicasetService) UpdateReplicaset(id string, n replicaset.ReplicasetDto) error {
	rsData, err0 := s.Repository.GetOneReplicaset(id)
	if err0 != nil {
		return fmt.Errorf("해당 데이터가 없음: %v", err0)
	}

	cmd := exec.Command("kubectl", "delete", "replicaset", rsData.MetadataName, "-n", "test") // test = namespace
	_, err2 := cmd.CombinedOutput()
	if err2 != nil {
		return fmt.Errorf("기존 replicaset 삭제실패: %v", err2)
	}

	replicasetManager := resource.New()
	err3 := replicasetManager.UserId("test").ApiVersion(n.ApiVersion).Kind(n.Kind).MetadataName(n.MetadataName).SpecReplicas(n.SpecReplicas).SpecSelectorMatchlabelsApp(n.SpecSelectorMatchlabelsApp).SpecTemplateMetadataName(n.SpecTemplateMetadataName).SpecTemplateMetadataLabelsApp(n.SpecTemplateMetadataLabelsApp).SpecTemplateSpecContainersName(n.SpecTemplateSpecContainersName).SpecTemplateSpecContainersImage(n.SpecTemplateSpecContainersImage).SpecTemplateSpecContainersPortsContainerport(n.SpecTemplateSpecContainersPortsContainerport).Build()
	if err3 != nil {
		return err3
	}
	_, err := s.Repository.UpdateOneReplicaset(id, n)
	if err != nil {
		return fmt.Errorf("db error: %v", err)
	}
	return nil
}

func (s *ReplicasetService) DeleteReplicaset(id string) error {
	rsData, err0 := s.Repository.GetOneReplicaset(id)
	if err0 != nil {
		return fmt.Errorf("해당 데이터가 없음: %v", err0)
	}

	cmd := exec.Command("kubectl", "delete", "replicaset", rsData.MetadataName, "-n", "test") // test = namesapce
	output, err2 := cmd.CombinedOutput()
	if err2 != nil {
		return fmt.Errorf("삭제실패: %v", err2)
	}
	_, err := s.Repository.DeleteOneReplicaset(id)
	if err != nil {
		fmt.Print(output)
		return fmt.Errorf("kubectl 명령 실행 중 오류 발생: %v", err2)

	}
	return err
}
func (s *ReplicasetService) GetReplicasetsStatus() (string, error) {
	// kubectl 명령 실행
	cmd := exec.Command("kubectl", "get", "replicasets", "-o", "json", "-n", "test") // test = namespace
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("kubectl 명령 실행 중 오류 발생: %v", err)
	}

	return string(output), nil
}
