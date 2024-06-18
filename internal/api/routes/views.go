package routes

import (
	"net/http"

	"github.com/covicale/url-shortener-go/internal/api/handlers"
)

func ViewsMux(handler *handlers.ViewHandler) http.Handler {
	viewsMux := http.NewServeMux()

	viewsMux.HandleFunc("GET /favicon.svg", handler.ServeFavicon)
	viewsMux.HandleFunc("GET /", handler.RenderHome)
	viewsMux.HandleFunc("GET /auth/login", handler.RenderLogin)
	viewsMux.HandleFunc("GET /auth/register", handler.RenderRegister)
	viewsMux.HandleFunc("GET /{url}", handler.RedirectToRealURL)

	return viewsMux
}
