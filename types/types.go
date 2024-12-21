package types

import "time"

type UserStore interface {
	GetUserByEmail(email string) (*User, error)

	// ?better like this? 
	//? CreateUser(user User) (*User, error)
	CreateUser(user User) error
}

type LoginUserPayload struct {
	Email 		string		`json:"email" validate:"required"`
	Password	string 		`json:"password" validate:"required"`
}

type RegisterPayload struct {
	Username	string	 	`json:"username" validate:"required"`
	Email 		string		`json:"email" validate:"required"`
	Password	string		`json:"password" validate:"required"`
}

type User struct {
	Id 			int
	Username	string
	Email		string
	Password	string
	CreatedAt	time.Time
}
