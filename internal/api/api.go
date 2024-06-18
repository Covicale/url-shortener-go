package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/covicale/url-shortener-go/internal/api/handlers"
	"github.com/covicale/url-shortener-go/internal/api/middleware"
	"github.com/covicale/url-shortener-go/internal/api/routes"
	"github.com/rs/cors"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	mux := http.NewServeMux()

	// Handlers
	urlHandler := handlers.NewURLHandler(s.db)
	authHandler := handlers.NewAuthHandler(s.db)
	viewHandler := handlers.NewViewHandler(s.db)

	// View routes
	mux.Handle("/", routes.ViewsMux(viewHandler))

	// API Routes grouped
	mux.Handle("/api/v1/url/", routes.URLMux(urlHandler))
	mux.Handle("/api/v1/auth/", routes.AuthMux(authHandler))

	// Middleware
	wrappedMux := middleware.NewAuthMiddleware(mux, s.db)

	// Cors
	corsHandler := cors.AllowAll().Handler(wrappedMux)

	log.Printf("App is running on port %v\n", s.addr)
	err := http.ListenAndServe(s.addr, corsHandler)
	return err
}
