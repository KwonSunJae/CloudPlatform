package reqchecker

import (
	"errors"
	"reflect"
)

// Check 함수는 들어온 generic 인자의 구성요소에 빈 값이 있는지 확인하고, 빈 값이 있으면 에러를 반환한다.
func Check[T any](r T) error {
	//r은 GENERIC
	//r의 타입이 struct인지 확인
	//r의 필드를 순회하면서 빈 값이 있는지 확인
	//빈 값이 있으면 에러 반환
	//모든 필드를 순회하고 빈 값이 없으면 nil 반환

	v := reflect.ValueOf(r)
	if v.Kind() != reflect.Struct {
		return errors.New("invalid request body")
	}

	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Interface() == "" {
			return errors.New("empty value")
		}
	}
	return nil
}
