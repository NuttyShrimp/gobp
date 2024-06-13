package dto

import "github.com/studentkickoff/gobp/pkg/sqlc"

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Uid   string `json:"uid"`
}

func UserDTO(user sqlc.User) User {
	return User{
		Name:  user.Name,
		Uid:   user.Uid,
		Email: user.Email,
	}
}
