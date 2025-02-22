package dto

import "github.com/studentkickoff/gobp/pkg/db/sqlc"

type User struct {
	ID    int    `json:"id"`
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

func (u *User) IntoCreateParams() sqlc.CreateUserParams {
	return sqlc.CreateUserParams{
		Name:  u.Name,
		Uid:   u.Uid,
		Email: u.Email,
	}
}
