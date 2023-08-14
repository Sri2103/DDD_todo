package user_model

import (
	"encoding/json"
	"io"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	UserName string    `json:"username"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

func NewUser() *User {
	return &User{}
}

func (u *User) ToJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(u)
}

func (u *User) FromJson(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(u)
}
