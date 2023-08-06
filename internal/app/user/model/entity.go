package user_model

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID
	Name     string
	UserName string
	Email    string
	Password string
}