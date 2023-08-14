package user_service

import (
	"github.com/google/uuid"
	user_model "github.com/sri2103/domain_DD_todo/internal/app/user/model"
	user_repository "github.com/sri2103/domain_DD_todo/internal/app/user/repository"
)

type UserServiceImpl struct {
	userRepo user_repository.UserRepository
}

func NewUserServiceImpl(userRepo user_repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{userRepo: userRepo}
}

func (s *UserServiceImpl) CreateUser(user *user_model.User) error {
	err := s.userRepo.Create(user)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserServiceImpl) GetUserById(id uuid.UUID) (*user_model.User, error) {
	user, err := s.userRepo.FindById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserServiceImpl) GetAllUsers() ([]*user_model.User, error) {
	users, err := s.userRepo.FindAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserServiceImpl) UpdateUser(user *user_model.User) error {
	err := s.userRepo.Update(user)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserServiceImpl) DeleteUser(id uuid.UUID) error {
	err := s.userRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
