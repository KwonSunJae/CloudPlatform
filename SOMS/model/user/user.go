package model

type User struct {
	Id          string
	Name        string
	UserID      string
	EncryptedPW string
	Role        string
	Spot        string
	Priority    string
}

type UserDto struct {
	Name        string
	UserID      string
	EncryptedPW string
	Role        string
	Spot        string
	Priority    string
}

type UserRaw struct {
	Id          string
	Name        string
	UserID      string
	EncryptedPW string
	Role        string
	Spot        string
	Priority    string
}
