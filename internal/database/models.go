package database

import "github.com/uptrace/bun"

type User struct {
	bun.BaseModel `bun:"table:users"`

	Id    int64 `bun:"pk,autoincrement"`
	Name  string
	Email string
	Uid   string
}
