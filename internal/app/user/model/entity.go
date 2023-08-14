package user_model

import (
	"encoding/json"
	"io"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID
	Name     string
	UserName string
	Email    string
	Password string
}

func NewUser() *User {
	return &User{}
}

func (u *User) ToJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(u)
}

func (u *User) fromJson(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(u)
}
