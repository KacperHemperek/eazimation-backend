package server

import (
	"eazimation-backend/internal/api"
	"eazimation-backend/internal/auth"
	"eazimation-backend/internal/handlers/health"
	"github.com/go-chi/cors"
	"github.com/gorilla/sessions"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (s *Server) RegisterRoutes(
	addProviderToCtx auth.AddProviderToContextMiddleware,
	sessionStore auth.SessionStore,
) http.Handler {
	r := chi.NewRouter()

	store := sessions.NewCookieStore([]byte("secret"))

	store.MaxAge(1000 * 60 * 60 * 24 * 7)
	store.Options.HttpOnly = true
	store.Options.Domain = "http://localhost:5173"
	store.Options.Secure = false

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{http.MethodPut, http.MethodPost, http.MethodGet, http.MethodDelete},
		AllowCredentials: true,
	}))

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/health", api.HttpHandler(health.GetHealthHandler()))

		r.Get("/auth/{provider}/callback", api.HttpHandler(addProviderToCtx(auth.HandleAuthCallback(sessionStore))))
		r.Get("/auth/{provider}", api.HttpHandler(addProviderToCtx(auth.HandleAuth(sessionStore))))
		r.Get("/logout/{provider}", api.HttpHandler(addProviderToCtx(auth.HandleLogout())))
		r.Get("/auth/user", api.HttpHandler(auth.HandleGetUser(sessionStore)))
	})

	return r
}
