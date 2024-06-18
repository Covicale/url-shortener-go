package routes

import (
	"net/http"

	"github.com/covicale/url-shortener-go/internal/api/handlers"
)

func AuthMux(handler *handlers.AuthHandler) http.Handler {
	authMux := http.NewServeMux()

	authMux.HandleFunc("POST /register", handler.RegisterUser)
	authMux.HandleFunc("POST /login", handler.LoginUser)
	authMux.HandleFunc("POST /logout", handler.LogoutUser)

	return http.StripPrefix("/api/v1/auth", authMux)
}
