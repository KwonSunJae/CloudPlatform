package resource

import (
	"fmt"
)

// 템플릿 관리하는 인터페이스
type TemplateManager interface {
	GetTemplate(string) (string, error)
}

// 템플릿 관리하는 클래스
type templateManager struct{}

// TemplateManager 인터페이스를 구현하는 새로운 객체 생성후 반환
// 반환값 templateManager 구조체의 인스턴스를 TemplateManager 인터페이스로 형변환
// 반환된 객체를 TemplateManager 타입으로 사용 가능
func NewTemplateManager() TemplateManager {
	return &templateManager{}
}

func (tm *templateManager) GetTemplate(specType string) (string, error) {
	switch specType {
	case "ClusterIP":
		return clusterIPTemplate, nil
	case "NodePort":
		return nodePortTemplate, nil
	case "LoadBalancer":
		return loadBalancerTemplate, nil
	case "ExternalName":
		return externalNameTemplate, nil
	default:
		return "", fmt.Errorf("지원되지 않는 SpecType입니다")
	}
}

const (
	clusterIPTemplate = `
apiVersion: {{.Dto.ApiVersion}}
kind: {{.Dto.Kind}}
metadata:
  name: {{.Dto.MetadataName}}
spec:
  type: {{.Dto.SpecType}}
  selector:
    app: {{.Dto.SpecSelectorApp}}
  ports:
    - protocol: {{.Dto.SpecPortsProtocol}}
      port: {{.Dto.SpecPortsPort}}
      targetPort: {{.Dto.SpecPortsTargetport}}
`

	nodePortTemplate = `
apiVersion: {{.Dto.ApiVersion}}
kind: {{.Dto.Kind}}
metadata:
  name: {{.Dto.MetadataName}}
spec:
  type: {{.Dto.SpecType}}
  selector:
    app: {{.Dto.SpecSelectorApp}}
  ports:
    - protocol: {{.Dto.SpecPortsProtocol}}
      port: {{.Dto.SpecPortsPort}}
      targetPort: {{.Dto.SpecPortsTargetport}}
      nodePort: {{.Dto.SpecPortsNodeport}}
`

	loadBalancerTemplate = `
apiVersion: {{.Dto.ApiVersion}}
kind: {{.Dto.Kind}}
metadata:
  name: {{.Dto.MetadataName}}
spec:
  type: {{.Dto.SpecType}}
  selector:
    app: {{.Dto.SpecSelectorApp}}
    type: {{.Dto.SpecSelectorType}}
  ports:
    - port: {{.Dto.SpecPortsProtocol}}
      protocol: {{.Dto.specPortsPort}}
      targetPort: {{.Dto.specPortsTargetport}}
  clusterIP: {{.Dto.SpecClusterIP}}
`

	externalNameTemplate = `
apiVersion: {{.Dto.ApiVersion}}
kind: {{.Dto.Kind}}
metadata:
  name: {{.Dto.MetadataName}}
  namespace: {{.Dto.MetadataNamespace}}
spec:
  type: {{.Dto.SpecType}}
  externalName: {{.Dto.SpecExternalname}}
`
)
