package utils

import (
	"errors"
	"net/http"
	"regexp"

	"github.com/covicale/url-shortener-go/internal/api/models"
	"github.com/covicale/url-shortener-go/internal/api/types"
	"github.com/covicale/url-shortener-go/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

func CreateJWTCookie(claims *types.UserJWTClaims) (*http.Cookie, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(config.Env.JWT_SECRET))
	if err != nil {
		return nil, err
	}

	return &http.Cookie{
		Name:     "jwt-access-token",
		Value:    tokenString,
		MaxAge:   3600, // 1 hour
		Path:     "/",
		HttpOnly: true,
	}, nil
}

func DeleteJWTCookie() *http.Cookie {
	return &http.Cookie{
		Name:     "jwt-access-token",
		Value:    "",
		MaxAge:   -1, // 1 hour
		Path:     "/",
		HttpOnly: true,
	}
}

func NewUserJWTClaims(user *models.User) *types.UserJWTClaims {
	return &types.UserJWTClaims{
		UserJWT: types.UserJWT{
			UserID:   user.Id,
			Username: user.Username,
		},
	}
}

func IsEmailValid(email string) bool {
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	return re.MatchString(email)
}

func IsURLValid(url string) bool {
	re := regexp.MustCompile(`^(https?|ftp):\/\/[^\s/$.?#].[^\s]*$`)
	return re.MatchString(url)
}

func CheckCredentialsOnRegister(credentials *types.RegisterUserRequest) error {
	// Check if the email is valid
	if !IsEmailValid(credentials.Email) {
		return errors.New("invalid email")
	}

	// Check if the password has between 4 and 16 characters
	if len(credentials.Password) < 4 || len(credentials.Password) > 16 {
		return errors.New("password must have between 4 and 16 characters")
	}

	// Check if the username has between 4 and 16 characters
	if len(credentials.Username) < 4 || len(credentials.Username) > 16 {
		return errors.New("username must have between 4 and 16 characters")
	}

	return nil
}
