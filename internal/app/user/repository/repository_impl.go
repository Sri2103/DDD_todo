package user_repository

import (
	"github.com/google/uuid"
	user_model "github.com/sri2103/domain_DD_todo/internal/app/user/model"
)

type UserRepositoryImpl struct {
	users []*user_model.User
}

func NewUserRepository()*UserRepositoryImpl{
	return &UserRepositoryImpl{users: []*user_model.User{}}
}

func (r *UserRepositoryImpl)Save(to *user_model.User) error {
	return nil
}


func (r *UserRepositoryImpl)FindById(id uuid.UUID)(*user_model.User,error){
	for _, user := range r.users{
		if user.ID == id {
			return user, nil
		}
	}
	return nil, nil
}

func (r *UserRepositoryImpl)FindAll()([]*user_model.User,error){
	return []*user_model.User{}, nil
}

func (r *UserRepositoryImpl)Update(user *user_model.User) error{
	return nil
}

func (r *UserRepositoryImpl)Delete(id uuid.UUID) error{
	return nil
}

