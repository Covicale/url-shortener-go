package middleware

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/covicale/url-shortener-go/internal/api/types"
	"github.com/covicale/url-shortener-go/internal/config"
	"github.com/covicale/url-shortener-go/internal/db/repositories"
	"github.com/golang-jwt/jwt/v5"
)

type AuthMiddleware struct {
	handler        http.Handler
	userRepository *repositories.UserRepository
}

func NewAuthMiddleware(handlerWrapped http.Handler, db *sql.DB) *AuthMiddleware {
	return &AuthMiddleware{
		handler:        handlerWrapped,
		userRepository: repositories.NewUserRepository(db),
	}
}

func (authMiddleware *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("jwt-access-token")
	if err != nil {
		authMiddleware.handler.ServeHTTP(w, r)
		return
	}

	jwtToken := cookie.Value
	token, err := jwt.ParseWithClaims(jwtToken, &types.UserJWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Env.JWT_SECRET), nil
	})

	if !token.Valid || err != nil {
		authMiddleware.handler.ServeHTTP(w, r)
		return
	}

	claims := token.Claims.(*types.UserJWTClaims)
	if userJson, err := json.Marshal(claims.UserJWT); err == nil {
		r.Header.Add("UserInfo", string(userJson))
	}
	authMiddleware.handler.ServeHTTP(w, r)
}
