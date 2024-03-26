package reqchecker

import (
	reqchecker "soms/controller/checker/user"
)

// Checker는 vm, user, container의 타입이 정해지면, 해당 타입에 맞는 Checker를 반환하는 함수를 정의합니다.
type Checker interface {
	SetChecker() Checker
	Check(body interface{}) error
}

// checker는 vm, user, container의 타입에 따라, 해당 타입에 맞는 Checker를 반환합니다.
type checker struct {
	CheckerType string
	Checker     interface{}
}

// New 함수는 Checker 인터페이스를 반환합니다.
func New(checkerType string) Checker {
	return &checker{
		CheckerType: checkerType,
	}
}

// GetChecker 함수는 checkerType에 따라, 해당 타입에 맞는 Checker를 반환합니다.
func (c *checker) SetChecker() Checker {
	switch c.CheckerType {
	case "user":
		c.Checker = reqchecker.UserRegisterChecker{}
		return c
	default:
		return nil
	}
}

func (c *checker) Check(r interface{}) error {
	switch c.CheckerType {
	case "user":
		return c.Checker.(reqchecker.UserRegisterChecker).Check(r.(reqchecker.RequestBody))
	default:
		return nil
	}
}
