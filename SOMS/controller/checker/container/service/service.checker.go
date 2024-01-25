package checker

import (
	"errors"
)

type RequestBody struct {
	ApiVersion          string
	Kind                string
	MetadataName        string
	SpecType            string
	SpecSelectorApp     string
	SpecPortsProtocol   string
	SpecPortsPort       string
	SpecPortsTargetport string
	SpecPortsNodeport   string
	SpecSelectorType    string
	SpecClusterIP       string
	MetadataNamespace   string
	SpecExternalname    string
}

func ServiceTypeChecker(body RequestBody) error {
	switch body.SpecType {
	case "ClusterIP":
		return checkClusterIP(body)
	case "NodePort":
		return checkNodePort(body)
	case "LoadBalancer":
		return checkLoadBalancer(body)
	case "ExternalName":
		return checkExternalName(body)
	default:
		return errors.New("유효하지 않은 SpecType입니다.")
	}
}

func checkClusterIP(body RequestBody) error {
	if body.ApiVersion == "" || body.Kind == "" || body.MetadataName == "" ||
		body.SpecType == "" || body.SpecSelectorApp == "" || body.SpecPortsProtocol == "" ||
		body.SpecPortsPort == "" || body.SpecPortsTargetport == "" {
		return errors.New("파라미터가 누락되었습니다.")
	}

	return nil
}

func checkNodePort(body RequestBody) error {
	if body.ApiVersion == "" || body.Kind == "" || body.MetadataName == "" ||
		body.SpecType == "" || body.SpecSelectorApp == "" || body.SpecPortsProtocol == "" ||
		body.SpecPortsPort == "" || body.SpecPortsTargetport == "" || body.SpecPortsNodeport == "" {
		return errors.New("파라미터가 누락되었습니다.")
	}

	return nil
}

func checkLoadBalancer(body RequestBody) error {
	if body.ApiVersion == "" || body.Kind == "" || body.MetadataName == "" ||
		body.SpecType == "" || body.SpecSelectorApp == "" || body.SpecSelectorType == "" ||
		body.SpecPortsProtocol == "" || body.SpecPortsPort == "" || body.SpecPortsTargetport == "" ||
		body.SpecClusterIP == "" {
		return errors.New("파라미터가 누락되었습니다.")
	}

	return nil
}

func checkExternalName(body RequestBody) error {
	if body.ApiVersion == "" || body.Kind == "" || body.MetadataName == "" ||
		body.MetadataNamespace == "" || body.SpecType == "" || body.SpecExternalname == "" {
		return errors.New("파라미터가 누락되었습니다.")
	}

	return nil
}
