package resource

import (
	"fmt"
	"os"
	"os/exec"
	"soms/repository/container/service"
	"text/template"
)

// Service 객체 생성 및 구성하는 추상화 인터페이스
type ServiceManager interface {
	UserID(string) ServiceManager
	ApiVersion(string) ServiceManager
	Kind(string) ServiceManager
	MetadataName(string) ServiceManager
	SpecType(string) ServiceManager
	SpecSelectorApp(string) ServiceManager
	SpecPortsProtocol(string) ServiceManager
	SpecPortsPort(string) ServiceManager
	SpecPortsTargetport(string) ServiceManager
	SpecPortsNodeport(string) ServiceManager
	SpecSelectorType(string) ServiceManager
	SpecClusterIP(string) ServiceManager
	SpecExternalname(string) ServiceManager
	Build() error
}

// 생성할 Service 객체
type serviceManager struct {
	Dto service.ServiceDto
}

// 새로운 ServiceManager 생성
func New() ServiceManager {
	return &serviceManager{}
}

func (sb *serviceManager) UserID(i string) ServiceManager {
	sb.Dto.UUID = i
	return sb
}

func (sb *serviceManager) ApiVersion(v string) ServiceManager {
	sb.Dto.ApiVersion = v
	return sb
}

func (sb *serviceManager) Kind(k string) ServiceManager {
	sb.Dto.Kind = k
	return sb
}

func (sb *serviceManager) MetadataName(mn string) ServiceManager {
	sb.Dto.MetadataName = mn
	return sb
}

func (sb *serviceManager) SpecType(t string) ServiceManager {
	sb.Dto.SpecType = t
	return sb
}

func (sb *serviceManager) SpecSelectorApp(sa string) ServiceManager {
	sb.Dto.SpecSelectorApp = sa
	return sb
}

func (sb *serviceManager) SpecPortsProtocol(pr string) ServiceManager {
	sb.Dto.SpecPortsProtocol = pr
	return sb
}

func (sb *serviceManager) SpecPortsPort(p string) ServiceManager {
	sb.Dto.SpecPortsPort = p
	return sb
}

func (sb *serviceManager) SpecPortsTargetport(tp string) ServiceManager {
	sb.Dto.SpecPortsTargetport = tp
	return sb
}

func (sb *serviceManager) SpecPortsNodeport(np string) ServiceManager {
	sb.Dto.SpecPortsNodeport = np
	return sb
}

func (sb *serviceManager) SpecSelectorType(st string) ServiceManager {
	sb.Dto.SpecSelectorType = st
	return sb
}

func (sb *serviceManager) SpecClusterIP(c string) ServiceManager {
	sb.Dto.SpecClusterIP = c
	return sb
}

func (sb *serviceManager) SpecExternalname(e string) ServiceManager {
	sb.Dto.SpecExternalname = e
	return sb
}

func (sb *serviceManager) Build() error {
	tm := NewTemplateManager()

	yamlTemplate, err := tm.GetTemplate(sb.Dto.SpecType)
	if err != nil {
		return err
	}

	// 템플릿에 데이터 적용
	tmpl, err := template.New("yaml").Parse(yamlTemplate)
	if err != nil {
		return fmt.Errorf("YAML 템플릿 파싱 중 오류 발생: %v", err)
	}

	fileName := fmt.Sprintf("k8s/%s/%s_service.yaml", sb.Dto.UUID, sb.Dto.MetadataName) // test = id
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
	err = tmpl.Execute(file, sb)
	if err != nil {
		return fmt.Errorf("YAML 파일 작성 중 오류 발생: %v", err)
	}

	// kubectl apply 실행
	cmd := exec.Command("kubectl", "apply", "-f", fileName, "-n", sb.Dto.UUID)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("kubectl apply 명령 실행 중 오류 발생: %v\nOutput: %s", err, output)
	}

	fmt.Printf("YAML 파일 생성/수정 및 kubectl apply 완료: %s\n", fileName)

	return nil
}
