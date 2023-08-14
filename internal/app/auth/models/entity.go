package auth_model

import "github.com/google/uuid"

type Auth_entity struct {
	ID       uuid.UUID
	Username string
	Password string
	email    string
}
