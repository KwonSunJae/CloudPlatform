package deployment

import (
	"fmt"
	"os"
	"os/exec"
	"soms/repository"
	"soms/repository/container/deployment"
	"text/template"
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

	templateStr := `
apiVersion: {{.ApiVersion}}
kind: {{.Kind}}
metadata:
  name: {{.Metadata_name}}
  labels:
    app: {{.Metadata_labels_app}}
spec:
  selector:
    matchLabels:
      app: {{.Spec_selector_matchLabels_app}}
  template:
    metadata:
      labels:
        app: {{.Spec_template_metadata_labels_app}}
    spec:
      hostname: {{.Spec_template_spec_hostname}}
      subdomain: {{.Spec_template_spec_subdomain}}
      containers:
      - image: {{.Spec_template_spec_containers_image}}
        imagePullPolicy: {{.Spec_template_spec_containers_imagePullPolicy}}
        name: {{.Spec_template_spec_containers_name}}
        ports:
        - containerPort: {{.Spec_template_spec_containers_ports_containerPort}}
`

	fileName := fmt.Sprintf("k8s/test/%s_deployment.yaml", n.Metadata_name)
	err1 := createYAMLFile(templateStr, fileName, n)
	if err1 != nil {
		return err1
	}
	_, err := s.Repository.InsertDeployment(n)
	if err != nil {
		return err
	}

	return applyKubectl(fileName)
}

func (s *DeploymentService) UpdateDeployment(id string, n deployment.DeploymentDto) error {
	_, err := s.Repository.UpdateOneDeployment(id, n)

	return err
}

func (s *DeploymentService) DeleteDeployment(id string) error {
	dpData, err0 := s.Repository.GetOneDeployment(id)
	if err0 != nil {
		return fmt.Errorf("해당 데이터가 없음: %v", err0)
	}

	cmd := exec.Command("kubectl", "delete", "deployment", dpData.Metadata_name)
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
	cmd := exec.Command("kubectl", "get", "deployments", "-o", "json")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("kubectl 명령 실행 중 오류 발생: %v", err)
	}

	return string(output), nil
}

func createYAMLFile(templateStr string, fileName string, data interface{}) error {
	tmpl, err := template.New("yaml").Parse(templateStr)
	if err != nil {
		return fmt.Errorf("YAML 템플릿 파싱 중 오류 발생: %v", err)
	}

	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("파일 생성 중 오류 발생: %v", err)
	}
	defer file.Close()

	err = tmpl.Execute(file, data)
	if err != nil {
		return fmt.Errorf("YAML 파일 작성 중 오류 발생: %v", err)
	}

	return nil
}

// applyKubectl 함수는 주어진 YAML 파일을 kubectl apply 명령으로 실행하는 함수입니다.
func applyKubectl(fileName string) error {
	cmd := exec.Command("kubectl", "apply", "-f", fileName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("kubectl apply 명령 실행 중 오류 발생: %v\nOutput: %s", err, output)
	}

	fmt.Printf("YAML 파일 생성 및 kubectl apply 완료: %s\n", fileName)
	return nil
}
