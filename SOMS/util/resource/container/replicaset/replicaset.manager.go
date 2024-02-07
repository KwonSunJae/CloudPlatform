package resource

import (
	"fmt"
	"os"
	"os/exec"
	"soms/repository/container/replicaset"
	"text/template"
)

// Replicaset 객체 생성 및 구성하는 추상화 인터페이스
type ReplicasetManager interface {
	UserId(string) ReplicasetManager
	ApiVersion(string) ReplicasetManager
	Kind(string) ReplicasetManager
	MetadataName(string) ReplicasetManager
	SpecReplicas(string) ReplicasetManager
	SpecSelectorMatchlabelsApp(string) ReplicasetManager
	SpecTemplateMetadataName(string) ReplicasetManager
	SpecTemplateMetadataLabelsApp(string) ReplicasetManager
	SpecTemplateSpecContainersName(string) ReplicasetManager
	SpecTemplateSpecContainersImage(string) ReplicasetManager
	SpecTemplateSpecContainersPortsContainerport(string) ReplicasetManager
	Build() error
}

// 생성할 Replicaset 객체
type replicasetManager struct {
	userId string
	Dto    replicaset.ReplicasetDto
}

// 새로운 ReplicasetManager 생성
func New() ReplicasetManager {
	return &replicasetManager{}
}

func (rm replicasetManager) UserId(i string) ReplicasetManager {
	rm.userId = i
	return rm
}

func (rm replicasetManager) ApiVersion(v string) ReplicasetManager {
	rm.Dto.ApiVersion = v
	return rm
}

func (rm replicasetManager) Kind(k string) ReplicasetManager {
	rm.Dto.Kind = k
	return rm
}

func (rm replicasetManager) MetadataName(mn string) ReplicasetManager {
	rm.Dto.MetadataName = mn
	return rm
}

func (rm replicasetManager) SpecReplicas(sr string) ReplicasetManager {
	rm.Dto.SpecReplicas = sr
	return rm
}

func (rm replicasetManager) SpecSelectorMatchlabelsApp(ma string) ReplicasetManager {
	rm.Dto.SpecSelectorMatchlabelsApp = ma
	return rm
}

func (rm replicasetManager) SpecTemplateMetadataName(la string) ReplicasetManager {
	rm.Dto.SpecTemplateMetadataName = la
	return rm
}

func (rm replicasetManager) SpecTemplateMetadataLabelsApp(mla string) ReplicasetManager {
	rm.Dto.SpecTemplateMetadataLabelsApp = mla
	return rm
}

func (rm replicasetManager) SpecTemplateSpecContainersName(cn string) ReplicasetManager {
	rm.Dto.SpecTemplateSpecContainersName = cn
	return rm
}

func (rm replicasetManager) SpecTemplateSpecContainersImage(ci string) ReplicasetManager {
	rm.Dto.SpecTemplateSpecContainersImage = ci
	return rm
}

func (rm replicasetManager) SpecTemplateSpecContainersPortsContainerport(cp string) ReplicasetManager {
	rm.Dto.SpecTemplateSpecContainersPortsContainerport = cp
	return rm
}

func (rm replicasetManager) Build() error {

	yamlTemplate := `
apiVersion: {{.Dto.ApiVersion}}
kind: {{.Dto.Kind}}
metadata:
  name: {{.Dto.MetadataName}}
spec:
  replicas: {{.Dto.SpecReplicas}}
  selector:
    matchLabels:
      app: {{.Dto.SpecSelectorMatchlabelsApp}}
  template:
    metadata:
      name: {{.Dto.SpecTemplateMetadataName}}
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

	fileName := fmt.Sprintf("k8s/test/%s_replicaset.yaml", rm.Dto.MetadataName) // test = id
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
	err = tmpl.Execute(file, rm)
	if err != nil {
		return fmt.Errorf("YAML 파일 작성 중 오류 발생: %v", err)
	}

	// kubectl apply 실행
	cmd := exec.Command("kubectl", "apply", "-f", fileName, "-n", "test") // test = namespace
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("kubectl apply 명령 실행 중 오류 발생: %v\nOutput: %s", err, output)
	}

	fmt.Printf("YAML 파일 생성/수정 및 kubectl apply 완료: %s\n", fileName)

	return nil
}
