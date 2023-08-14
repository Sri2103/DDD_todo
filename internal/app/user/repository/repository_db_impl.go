package user_repository

import (
	"github.com/google/uuid"
	user_model "github.com/sri2103/domain_DD_todo/internal/app/user/model"
	"gorm.io/gorm"
)

type UserPostgresImpl struct {
	db *gorm.DB
}

type User_Pg_Todo struct {
	*gorm.Model
	ID       uuid.UUID `gorm:"primarykey"`
	Name     string
	UserName string
	Email    string
	Password string
}

func NewUserPostgresImpl(db *gorm.DB) *UserPostgresImpl {
	return &UserPostgresImpl{
		db: db,
	}
}

func (r *UserPostgresImpl) Create(user *user_model.User) error {
	return r.db.Create(&User_Pg_Todo{
		ID:       uuid.New(),
		Name:     user.Name,
		UserName: user.UserName,
		Email:    user.Email,
		Password: user.Password,
	}).Error
}

func (r *UserPostgresImpl) FindById(id uuid.UUID) (*user_model.User, error) {
	var user = new(User_Pg_Todo)
	if err := r.db.Where("id=?", id).First(&user).Error; err != nil {
		return nil, err
	}
	// return &user_model.User{
	// 	ID:       id,
	// 	Name:     user.Name,
	// 	UserName: user.UserName,
	// 	Email:    user.Email,
	// 	Password: user.Password,
	// }, nil

	return ConvertToEntity(user)

}

func (r *UserPostgresImpl) FindAll() ([]*user_model.User, error) {
	var Pg_users []*User_Pg_Todo
	if err := r.db.Find(&Pg_users).Error; err != nil {
		return nil, err
	}

	var users []*user_model.User

	for _, u := range Pg_users {
		user, err := ConvertToEntity(u)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *UserPostgresImpl) Update(user *user_model.User) error {
	// var user_Pg = &User_Pg_Todo{
	// 	ID:       user.ID,
	// 	Name:     user.Name,
	// 	UserName: user.UserName,
	// 	Email:    user.Email,
	// }
	user_Pg,err := ConvertToDbSchema(user)
	if err != nil {
		return nil
	}
	err = r.db.Save(user_Pg).Error //update

	if err != nil {
		return nil
	}

	return nil

}

func (r *UserPostgresImpl) Delete(id uuid.UUID) error {
	var user User_Pg_Todo
	//delete from table where
	err := r.db.Delete(&user, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserPostgresImpl) FindByUserNameAndPassword(userName, password string) (*user_model.User, error) {
	var user *User_Pg_Todo
	err := r.db.First(user, userName, password).Error

	if err != nil {
		return nil, err
	}
	return ConvertToEntity(user)
}
