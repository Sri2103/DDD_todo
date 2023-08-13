package user_repository

import (
	"errors"

	"github.com/google/uuid"
	user_model "github.com/sri2103/domain_DD_todo/internal/app/user/model"
)

type UserRepository interface {
	Create(user *user_model.User) error
	FindById(id uuid.UUID) (*user_model.User, error)
	FindAll() ([]*user_model.User, error)
	Update(user *user_model.User) error
	Delete(id uuid.UUID) error
	FindByUserNameAndPassword(userName, password string) (*user_model.User, error)
}

func ConvertToEntity(user *User_Pg_Todo) (*user_model.User, error) {
	if user.ID == uuid.Nil {
		return nil, errors.New("empty id")
	}
	var entity = &user_model.User{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		UserName: user.UserName,
	}
	return entity, nil
}

func ConvertToDbSchema(user *user_model.User) (*User_Pg_Todo, error) {
	if user.ID == uuid.Nil {
		return nil, errors.New("empty id")
	}
	var entity = &User_Pg_Todo{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		UserName: user.UserName,
		Password: user.Password,
	}
	return entity, nil
}
