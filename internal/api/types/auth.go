package types

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type UserJWTClaims struct {
	UserJWT
	jwt.RegisteredClaims
}

type UserJWT struct {
	UserID   uuid.UUID `json:"userID"`
	Username string    `json:"username"`
}

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}
