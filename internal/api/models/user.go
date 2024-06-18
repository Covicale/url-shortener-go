package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

func NewUser(username, password, email string) *User {
	return &User{
		Id:        uuid.New(),
		Username:  username,
		Password:  password,
		Email:     email,
		CreatedAt: time.Now(),
	}
}
