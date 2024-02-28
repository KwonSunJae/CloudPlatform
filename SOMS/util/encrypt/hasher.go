package encrypt

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// PasswordHasher 인터페이스 정의
type PasswordHasher interface {
	HashPassword(password string) (string, error)
	ComparePassword(plainPassword, hashedPassword string) error
}

// BcryptPasswordHasher 구조체 정의
type BcryptPasswordHasher struct {
	// SECRET 키를 저장하는 필드 추가
	SecretKey string
}

// NewBcryptPasswordHasher 생성자 함수 정의
func NewPasswordHasher(secretKey string) *BcryptPasswordHasher {

	return &BcryptPasswordHasher{
		SecretKey: secretKey,
	}
}

// HashPassword 메서드 구현
func (b *BcryptPasswordHasher) HashPassword(password string) (string, error) {
	fmt.Print(b.SecretKey)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password+b.SecretKey), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// cmp 구현
func (b *BcryptPasswordHasher) ComparePassword(plainPassword, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
}
