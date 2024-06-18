package routes

import (
	"net/http"

	"github.com/covicale/url-shortener-go/internal/api/handlers"
)

func URLMux(handler *handlers.URLHandler) http.Handler {
	urlMux := http.NewServeMux()

	urlMux.HandleFunc("POST /create", handler.CreateURL)
	urlMux.HandleFunc("DELETE /{shortURL}", handler.DeleteURL)

	return http.StripPrefix("/api/v1/url", urlMux)
}
