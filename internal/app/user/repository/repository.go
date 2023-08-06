package user_repository

import (
	"github.com/google/uuid"
	user_model "github.com/sri2103/domain_DD_todo/internal/app/user/model"
)

type UserRepository interface {
	Save(user *user_model.User) error
	FindById(id uuid.UUID)(*user_model.User,error)
	FindAll()([]*user_model.User,error)
	Update(user *user_model.User) error
	Delete(id uuid.UUID) error
}