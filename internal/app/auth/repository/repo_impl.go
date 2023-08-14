package auth_repository

import (
	"github.com/google/uuid"
	auth_model "github.com/sri2103/domain_DD_todo/internal/app/auth/models"
	"gorm.io/gorm"
)

type Auth_entity_Pg struct {
	ID       uuid.UUID
	Username string
	Password string
	Email    string
	*gorm.Model
}
type AuthRepositoryImpl struct {
	db *gorm.DB
}

func NewAuthRepositoryImpl(db *gorm.DB) *AuthRepositoryImpl {
	return &AuthRepositoryImpl{db: db}
}

func (r *AuthRepositoryImpl) FindByUserNameAndPassword(username, password string) (*auth_model.Auth_entity, error) {
	var authEntity Auth_entity_Pg
	if err := r.db.Where("username=?", username).First(&authEntity).Error; err != nil {
		return nil, err
	}
	if authEntity.Password != password {
		return nil, gorm.ErrRecordNotFound
	}
	return &auth_model.Auth_entity{
		ID:       authEntity.ID,
		Username: authEntity.Username,
		Password: authEntity.Password,
	}, nil
}
