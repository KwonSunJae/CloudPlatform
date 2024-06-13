package resource

import (
	"fmt"
	"os"
	"os/exec"
	"soms/repository/container/deployment"
	"text/template"
)

// Deployment 객체 생성 및 구성하는 추상화 인터페이스
type DeploymentManager interface {
	UserID(string) DeploymentManager
	ApiVersion(string) DeploymentManager
	Kind(string) DeploymentManager
	MetadataName(string) DeploymentManager
	MetadataLabelsApp(string) DeploymentManager
	SpecReplicas(string) DeploymentManager
	SpecSelectorMatchlabelsApp(string) DeploymentManager
	SpecTemplateMetadataLabelsApp(string) DeploymentManager
	SpecTemplateSpecContainersName(string) DeploymentManager
	SpecTemplateSpecContainersImage(string) DeploymentManager
	SpecTemplateSpecContainersPortsContainerport(string) DeploymentManager
	Build() error
}

// 생성할 Deployment 객체
type deploymentManager struct {
	Dto deployment.DeploymentDto
}

// 새로운 DeploymentManager 생성
func New() DeploymentManager {
	return &deploymentManager{}
}

func (dm deploymentManager) UserID(i string) DeploymentManager {
	dm.Dto.UUID = i
	return dm
}

func (dm deploymentManager) ApiVersion(v string) DeploymentManager {
	dm.Dto.ApiVersion = v
	return dm
}

func (dm deploymentManager) Kind(k string) DeploymentManager {
	dm.Dto.Kind = k
	return dm
}

func (dm deploymentManager) MetadataName(mn string) DeploymentManager {
	dm.Dto.MetadataName = mn
	return dm
}

func (dm deploymentManager) MetadataLabelsApp(la string) DeploymentManager {
	dm.Dto.MetadataLabelsApp = la
	return dm
}

func (dm deploymentManager) SpecReplicas(sr string) DeploymentManager {
	dm.Dto.SpecReplicas = sr
	return dm
}

func (dm deploymentManager) SpecSelectorMatchlabelsApp(ma string) DeploymentManager {
	dm.Dto.SpecSelectorMatchlabelsApp = ma
	return dm
}

func (dm deploymentManager) SpecTemplateMetadataLabelsApp(mla string) DeploymentManager {
	dm.Dto.SpecTemplateMetadataLabelsApp = mla
	return dm
}

func (dm deploymentManager) SpecTemplateSpecContainersName(cn string) DeploymentManager {
	dm.Dto.SpecTemplateSpecContainersName = cn
	return dm
}

func (dm deploymentManager) SpecTemplateSpecContainersImage(ci string) DeploymentManager {
	dm.Dto.SpecTemplateSpecContainersImage = ci
	return dm
}

func (dm deploymentManager) SpecTemplateSpecContainersPortsContainerport(cp string) DeploymentManager {
	dm.Dto.SpecTemplateSpecContainersPortsContainerport = cp
	return dm
}

func (dm deploymentManager) Build() error {

	yamlTemplate := `
apiVersion: {{.Dto.ApiVersion}}
kind: {{.Dto.Kind}}
metadata:
  name: {{.Dto.MetadataName}}
  labels:
    app: {{.Dto.MetadataLabelsApp}}
spec:
  replicas: {{.Dto.SpecReplicas}}
  selector:
    matchLabels:
      app: {{.Dto.SpecSelectorMatchlabelsApp}}
  template:
    metadata:
      labels:
        app: {{.Dto.SpecTemplateMetadataLabelsApp}}
    spec:
      containers:
        - name: {{.Dto.SpecTemplateSpecContainersName}}
          image: {{.Dto.SpecTemplateSpecContainersImage}}
          ports:
            - containerPort: {{.Dto.SpecTemplateSpecContainersPortsContainerport}}
`

	// 템플릿에 데이터 적용
	tmpl, err := template.New("yaml").Parse(yamlTemplate)
	if err != nil {
		return fmt.Errorf("YAML 템플릿 파싱 중 오류 발생: %v", err)
	}

	fileName := fmt.Sprintf("k8s/%s/%s_deployment.yaml", dm.Dto.UUID, dm.Dto.MetadataName) // test = id
	var file *os.File

	if _, err = os.Stat(fileName); os.IsNotExist(err) {
		// 파일이 존재하지 않는 경우, 파일 생성
		file, err = os.Create(fileName)
		if err != nil {
			return fmt.Errorf("파일 생성 중 오류 발생: %v", err)
		}
		defer file.Close()
	} else {
		// 파일이 이미 존재하는 경우, 파일 불러오기
		file, err = os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			return fmt.Errorf("파일 불러오기 중 오류 발생: %v", err)
		}
		defer file.Close()
	}

	// 템플릿에 데이터 적용하여 파일에 쓰기
	err = tmpl.Execute(file, dm)
	if err != nil {
		return fmt.Errorf("YAML 파일 작성 중 오류 발생: %v", err)
	}

	// kubectl apply 실행
	cmd := exec.Command("kubectl", "apply", "-f", fileName, "-n", dm.Dto.UUID) // , "-n", dm.Dto.UserID
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("kubectl apply 명령 실행 중 오류 발생: %v\nOutput: %s", err, output)
	}

	fmt.Printf("YAML 파일 생성/수정 및 kubectl apply 완료: %s\n", fileName)

	return nil
}
