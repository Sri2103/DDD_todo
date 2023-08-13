package auth_repository

import auth_model "github.com/sri2103/domain_DD_todo/internal/app/auth/models"

type AuthRepository interface {
	FindByUserNameAndPassword(username, password string) (*auth_model.Auth_entity,error)
}