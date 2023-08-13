package auth_service

import (
	"errors"

	"github.com/sri2103/domain_DD_todo/internal/app/auth"
	auth_repository "github.com/sri2103/domain_DD_todo/internal/app/auth/repository"
)

type AuthServiceImplDB struct {
	authRepo auth_repository.AuthRepository
}

func (s *AuthServiceImplDB)Login(username,password string)(string,error){

	// Validate username and password
	authModel,err := s.authRepo.FindByUserNameAndPassword(username,password)
	if err!= nil{
		return "",errors.New("Invalid credentials")
	}

	token,err := auth.GenerateToken(authModel.ID.String(),authModel.Username)

	if err != nil {
		return "",err
	}

	return token, nil

}