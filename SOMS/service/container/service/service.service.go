package service

import (
	"fmt"
	"os"
	"os/exec"
	"soms/repository"
	"soms/repository/container/service"
	"text/template"
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

	yamlTemplate := `
apiVersion: {{.ApiVersion}}
kind: {{.Kind}}
metadata:
  name: {{.Metadata_name}}
spec:
  ports:
    - port: {{.Spec_ports_port}}
      protocol: {{.Spec_ports_protocol}}
      targetPort: {{.Spec_ports_targetPort}}
  selector:
    app: {{.Spec_selector_app}}
`

	// 템플릿에 데이터 적용
	tmpl, err := template.New("yaml").Parse(yamlTemplate)
	if err != nil {
		return fmt.Errorf("YAML 템플릿 파싱 중 오류 발생: %v", err)
	}

	// 파일 생성
	fileName := fmt.Sprintf("k8s/test/%s_service.yaml", n.Metadata_name)
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("파일 생성 중 오류 발생: %v", err)
	}
	defer file.Close()

	// 템플릿에 데이터 적용하여 파일에 쓰기
	err = tmpl.Execute(file, n)
	if err != nil {
		return fmt.Errorf("YAML 파일 작성 중 오류 발생: %v", err)
	}

	// kubectl apply 실행
	cmd := exec.Command("kubectl", "apply", "-f", fileName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("kubectl apply 명령 실행 중 오류 발생: %v\nOutput: %s", err, output)
	}
	_, err2 := s.Repository.InsertService(n)
	if err2 != nil {
		return fmt.Errorf("db error: %v\n", err2)
	}
	fmt.Printf("YAML 파일 생성 및 kubectl apply 완료: %s\n", fileName)
	return err
}

func (s *ServiceService) UpdateService(id string, n service.ServiceDto) error {
	// db에서 id에 해당하는 서비스를 새로운 serviceDTO로 업데이트
	_, err := s.Repository.UpdateOneService(id, n)
	if err != nil {
		return fmt.Errorf("db error : %v\n", err)
	}
	
	// 실행중인 yaml 파일을 불러와 새로운 DTO값으로 다시 작성 후 실행
	yamlTemplate := `
apiVersion: {{.ApiVersion}}
kind: {{.Kind}}
metadata:
  name: {{.Metadata_name}}
spec:
  ports:
    - port: {{.Spec_ports_port}}
      protocol: {{.Spec_ports_protocol}}
      targetPort: {{.Spec_ports_targetPort}}
  selector:
    app: {{.Spec_selector_app}}
`		

	// 템플릿에 데이터 적용
	tmpl, err := template.New("yaml").Parse(yamlTemplate)
	if err != nil {
		return fmt.Errorf("YAML 템플릿 파싱 중 오류 발생: %v", err)
	}

	// 파일 불러오기
	fileName := fmt.Sprintf("k8s/test/%s_service.yaml", n.Metadata_name)
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("파일 불러오기 중 오류 발생: %v", err)
	}
	defer file.Close()

	// 존재하는 YAML 파일의 내용 구조체로 읽어오기
	var existingServiceDto servcie.ServiceDto
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&existingServiceDto)
	if err != nil {
		return fmt.Errorf("YAML 파일 읽어오기 중 오류 발생: %v", err)
	}

	// 새로운 DTO 값과 기존 파일 내용 비교 및 수정
	if n.ApiVersion != existingServiceDto.ApiVersion {
		existingServiceDto.ApiVersion = n.ApiVersion
	}
	if n.Kind != existingServiceDto.Kind {
		existingServiceDto.Kind = n.Kind
	}
	if n.Metadata_name != existingServiceDto.Metadata_name {
		existingServiceDto.Metadata_name = n.Metadata_name
	}
	if n.Spec_ports_port != existingServiceDto.Spec_ports_port {
		existingServiceDto.Spec_ports_port = n.Spec_ports_port
	}
	if n.Spec_ports_protocol != existingServiceDto.Spec_ports_protocol {
		existingServiceDto.Spec_ports_protocol = n.Spec_ports_protocol
	}
	if n.Spec_ports_targetPort != existingServiceDto.Spec_ports_targetPort {
		existingServiceDto.Spec_ports_targetPort = n.Spec_ports_targetPort
	}
	if n.Spec_selector_app != existingServiceDto.Spec_selector_app {
		existingServiceDto.Spec_selector_app = n.Spec_selector_app
	}

	// 파일 업데이트
	file, err = os.Create(fileName)
	if err != nil {
		return fmt.Errorf("파일 업데이트 중 오류 발생: %v", err)
	}
	defer file.Close()

	// 수정된 내용을 템플릿에 적용하여 파일에 쓰기
	err = tmpl.Execute(file, existingServiceDto)
	if err != nil {
		return fmt.Errorf("YAML 파일 작성 중 오류 발생: %v", err)
	}

	// kubectl apply 실행
	cmd := exec.Command("kubectl", "apply", "-f", fileName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("kubectl apply 명령 실행 중 오류 발생: %v\nOutput: %s", err, output)
	}

	fmt.Printf("YAML 파일 수정 및 kubectl apply 완료: %s\n", fileName)
	return err
}

func (s *ServiceService) DeleteService(id string) error {
	svData, err0 := s.Repository.GetOneService(id)
	if err0 != nil {
		return fmt.Errorf("해당 데이터가 없음: %v", err0)
	}

	cmd := exec.Command("kubectl", "delete", "service", svData.Metadata_name)
	output, err2 := cmd.CombinedOutput()
	_, err := s.Repository.DeleteOneService(id)
	if err != nil {
		fmt.Print(output)
		return fmt.Errorf("kubectl 명령 실행 중 오류 발생: %v", err2)

	}
	return err
}

func (s *ServiceService) GetServiceStatus() (string, error) {
	// kubectl 명령 실행
	cmd := exec.Command("kubectl", "get", "services", "-o", "json")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("kubectl 명령 실행 중 오류 발생: %v", err)
	}

	return string(output), nil
}
