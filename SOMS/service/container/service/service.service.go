package service

import (
	"fmt"
	"os/exec"
	"soms/repository"
	"soms/repository/container/service"
	"soms/repository/user"
	resource "soms/util/resource/container/service"
)

type ServiceService struct {
	Repository *service.ServiceRepository
}

var Service ServiceService

func (s *ServiceService) InitService() error {
	db, err := repository.OpenWithFile()

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

	_, err2 := s.Repository.InsertService(n)
	if err2 != nil {
		return fmt.Errorf("db error: %v", err2)
	}

	return nil
}

func (s *ServiceService) ApproveService(id string) error {
	n, err := s.Repository.GetOneService(id)
	if err != nil {
		return fmt.Errorf("해당 데이터가 없음: %v", err)
	}
	targetUser, err := user.Repository.GetOneUserByUUID(n.UUID)
	if err != nil {
		return fmt.Errorf("해당 데이터가 없음: %v", err)
	}

	serviceManager := resource.New()
	if n.SpecType == "ClusterIP" {
		err := serviceManager.UserID(targetUser.UserID).ApiVersion(n.ApiVersion).Kind(n.Kind).MetadataName(n.MetadataName).SpecType(n.SpecType).SpecSelectorApp(n.SpecSelectorApp).SpecPortsProtocol(n.SpecPortsProtocol).SpecPortsPort(n.SpecPortsPort).SpecPortsTargetport(n.SpecPortsTargetport).Build()
		if err != nil {
			return err
		}
	}
	if n.SpecType == "NodePort" {
		err := serviceManager.UserID(targetUser.UserID).ApiVersion(n.ApiVersion).Kind(n.Kind).MetadataName(n.MetadataName).SpecType(n.SpecType).SpecSelectorApp(n.SpecSelectorApp).SpecPortsProtocol(n.SpecPortsProtocol).SpecPortsPort(n.SpecPortsPort).SpecPortsTargetport(n.SpecPortsTargetport).SpecPortsNodeport(n.SpecPortsNodeport).Build()
		if err != nil {
			return err
		}
	}
	if n.SpecType == "LoadBalancer" {
		err := serviceManager.UserID(targetUser.UserID).ApiVersion(n.ApiVersion).Kind(n.Kind).MetadataName(n.MetadataName).SpecType(n.SpecType).SpecSelectorApp(n.SpecSelectorApp).SpecSelectorType(n.SpecSelectorType).SpecPortsProtocol(n.SpecPortsProtocol).SpecPortsPort(n.SpecPortsPort).SpecPortsTargetport(n.SpecPortsTargetport).SpecClusterIP(n.SpecClusterIP).Build()
		if err != nil {
			return err
		}
	}
	if n.SpecType == "ExternalName" {
		err := serviceManager.UserID(targetUser.UserID).ApiVersion(n.ApiVersion).Kind(n.Kind).MetadataName(n.MetadataName).SpecType(n.SpecType).SpecExternalname(n.SpecExternalname).Build()
		if err != nil {
			return err
		}
	}

	var approvedService service.ServiceDto
	approvedService.Status = "Approved"

	_, err2 := s.Repository.UpdateOneService(id, approvedService)
	if err2 != nil {
		return fmt.Errorf("db error: %v", err2)
	}

	return nil
}

func (s *ServiceService) UpdateService(id string, n service.ServiceDto) error {
	svData, err0 := s.Repository.GetOneService(id)
	if err0 != nil {
		return fmt.Errorf("해당 데이터가 없음: %v", err0)
	}
	targetUser, err := user.Repository.GetOneUserByUUID(svData.UUID)
	if err != nil {
		return fmt.Errorf("해당 데이터가 없음: %v", err)
	}

	cmd := exec.Command("kubectl", "delete", "service", svData.MetadataName, "-n", targetUser.UserID)
	_, err2 := cmd.CombinedOutput()
	if err2 != nil {
		return fmt.Errorf("기존 service 삭제실패: %v", err2)
	}

	serviceManager := resource.New()
	if n.SpecType == "ClusterIP" {
		err := serviceManager.UserID(targetUser.UserID).ApiVersion(n.ApiVersion).Kind(n.Kind).MetadataName(n.MetadataName).SpecType(n.SpecType).SpecSelectorApp(n.SpecSelectorApp).SpecPortsProtocol(n.SpecPortsProtocol).SpecPortsPort(n.SpecPortsPort).SpecPortsTargetport(n.SpecPortsTargetport).Build()
		if err != nil {
			return err
		}
	}
	if n.SpecType == "NodePort" {
		err := serviceManager.UserID(targetUser.UserID).ApiVersion(n.ApiVersion).Kind(n.Kind).MetadataName(n.MetadataName).SpecType(n.SpecType).SpecSelectorApp(n.SpecSelectorApp).SpecPortsProtocol(n.SpecPortsProtocol).SpecPortsPort(n.SpecPortsPort).SpecPortsTargetport(n.SpecPortsTargetport).SpecPortsNodeport(n.SpecPortsNodeport).Build()
		if err != nil {
			return err
		}
	}
	if n.SpecType == "LoadBalancer" {
		err := serviceManager.UserID(targetUser.UserID).ApiVersion(n.ApiVersion).Kind(n.Kind).MetadataName(n.MetadataName).SpecType(n.SpecType).SpecSelectorApp(n.SpecSelectorApp).SpecSelectorType(n.SpecSelectorType).SpecPortsProtocol(n.SpecPortsProtocol).SpecPortsPort(n.SpecPortsPort).SpecPortsTargetport(n.SpecPortsTargetport).SpecClusterIP(n.SpecClusterIP).Build()
		if err != nil {
			return err
		}
	}
	if n.SpecType == "ExternalName" {
		err := serviceManager.UserID(targetUser.UserID).ApiVersion(n.ApiVersion).Kind(n.Kind).MetadataName(n.MetadataName).SpecType(n.SpecType).SpecExternalname(n.SpecExternalname).Build()
		if err != nil {
			return err
		}
	}

	_, err3 := s.Repository.UpdateOneService(id, n)
	if err3 != nil {
		return fmt.Errorf("db error : %v", err)
	}

	return nil
}

func (s *ServiceService) DeleteService(id string) error {
	svData, err0 := s.Repository.GetOneService(id)
	if err0 != nil {
		return fmt.Errorf("해당 데이터가 없음: %v", err0)
	}
	targetUser, err := user.Repository.GetOneUserByUUID(svData.UUID)
	if err != nil {
		return fmt.Errorf("해당 데이터가 없음: %v", err)
	}

	cmd := exec.Command("kubectl", "delete", "service", svData.MetadataName, "-n", targetUser.UserID)
	output, err2 := cmd.CombinedOutput()
	_, err3 := s.Repository.DeleteOneService(id)
	if err3 != nil {
		fmt.Print(output)
		return fmt.Errorf("kubectl 명령 실행 중 오류 발생: %v", err2)

	}
	return err
}

func (s *ServiceService) GetServiceStatus(uuid string) (string, error) {
	targetUser, err := user.Repository.GetOneUserByUUID(uuid)
	if err != nil {
		return "", fmt.Errorf("해당 데이터가 없음: %v", err)
	}
	// kubectl 명령 실행
	cmd := exec.Command("kubectl", "get", "services", "-o", "json", "-n", targetUser.UserID) // 실행중인 서비스 정보 json으로 출력
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("kubectl 명령 실행 중 오류 발생: %v", err)
	}

	return string(output), nil
}
