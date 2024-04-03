package server

import (
	"eazimation-backend/internal/api"
	"eazimation-backend/internal/handlers/health"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/health", api.HttpHandler(health.GetHealthHandler()))
	})

	return r
}
