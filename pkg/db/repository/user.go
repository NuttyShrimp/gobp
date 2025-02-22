package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/studentkickoff/gobp/internal/database/dto"
	"github.com/studentkickoff/gobp/pkg/db"
)

type User interface {
	GetById(context.Context, int32) (*dto.User, error)
	GetByUid(context.Context, string) (*dto.User, error)
	Create(context.Context, *dto.User) error
}

type userRepo struct {
	db db.DB
}

var _ User = (*userRepo)(nil)

func (r *userRepo) GetById(c context.Context, id int32) (*dto.User, error) {
	dbUser, err := r.db.Queries().GetUser(c, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, err
		} else {
			defaultUser := dto.User{}
			return &defaultUser, nil
		}
	}

	dtoUser := dto.UserDTO(dbUser)
	return &dtoUser, nil
}

func (r *userRepo) GetByUid(c context.Context, uid string) (*dto.User, error) {
	dbUser, err := r.db.Queries().GetUserByUid(c, uid)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, err
		} else {
			defaultUser := dto.User{}
			return &defaultUser, nil
		}
	}

	user := dto.UserDTO(dbUser)
	return &user, nil
}

func (r *userRepo) Create(c context.Context, user *dto.User) error {
	dbUser, err := r.db.Queries().CreateUser(c, user.IntoCreateParams())
	if err != nil {
		return err
	}

	user.ID = int(dbUser.ID)

	return nil
}
