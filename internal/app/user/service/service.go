package user_service

import (
	"github.com/google/uuid"
	user_model "github.com/sri2103/domain_DD_todo/internal/app/user/model"
)

type UserService interface {
	CreateUser(user *user_model.User) error
	GetUserById(id uuid.UUID) (*user_model.User, error)
	GetAllUsers() ([]*user_model.User, error)
	UpdateUser(user *user_model.User) error
	DeleteUser(id uuid.UUID) error
}
