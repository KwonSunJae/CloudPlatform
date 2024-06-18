package userchecker

// RequestBody는 User 정보를 담고 있는 요청 본문입니다.
type UserRegisterRequestBody struct {
	Name     string
	UserID   string
	PW       string
	Role     string
	Spot     string
	Priority string
}

type UserLoginRequestBody struct {
	UserID string
	PW     string
}
