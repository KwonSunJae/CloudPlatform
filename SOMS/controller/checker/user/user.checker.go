package reqchecker

import (
	"errors"
)

// RequestBody는 User 정보를 담고 있는 요청 본문입니다.
type RequestBody struct {
	Name     string
	UserID   string
	PW       string
	Role     string
	Spot     string
	Priority string
}

// UserRequestBodyChecker 인터페이스는 User의 RequestBody를 검사하는 함수를 정의합니다.
type UserRequestBodyChecker interface {
	Check(r RequestBody) error
}

// UserRegisterChecker 구조체는 UserRequestBodyChecker 인터페이스를 구현합니다.
type UserRegisterChecker struct{}

// UserRegisterCheck 함수는 UserRegisterChecker의 메서드로서, UserRequestBody를 검사합니다.
func (c UserRegisterChecker) Check(r RequestBody) error {
	if r.Name == "" || r.UserID == "" || r.PW == "" || r.Role == "" || r.Spot == "" || r.Priority == "" {
		return errors.New("파라미터가 누락되었습니다.")
	}
	return nil
}
