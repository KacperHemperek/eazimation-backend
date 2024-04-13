package server

import (
	"eazimation-backend/internal/api"
	"eazimation-backend/internal/auth"
	"eazimation-backend/internal/handlers/health"
	"github.com/go-chi/cors"
	"github.com/gorilla/sessions"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

var (
	frontendUrl   = os.Getenv("FRONTEND_URL")
	sessionSecret = os.Getenv("SESSION_SECRET")
)

func (s *Server) RegisterRoutes(
	addProviderToCtx auth.AddProviderToContextMiddleware,
	sessionStore auth.SessionStore,
) http.Handler {
	r := chi.NewRouter()

	store := sessions.NewCookieStore([]byte(sessionSecret))

	store.MaxAge(1000 * 60 * 60 * 24 * 7)
	store.Options.HttpOnly = true
	store.Options.Secure = false

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{frontendUrl},
		AllowedMethods:   []string{http.MethodPut, http.MethodPost, http.MethodGet, http.MethodDelete},
		AllowCredentials: true,
	}))

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/health", api.HttpHandler(health.GetHealthHandler()))

		r.Get("/auth/{provider}/callback", api.HttpHandler(addProviderToCtx(auth.HandleAuthCallback(sessionStore))))
		r.Get("/auth/{provider}", api.HttpHandler(addProviderToCtx(auth.HandleAuth(sessionStore))))
		r.Post("/auth/logout", api.HttpHandler(auth.HandleLogout(sessionStore)))
		r.Get("/auth/user", api.HttpHandler(auth.HandleGetUser(sessionStore)))
		r.Get("/auth/lambda", api.HttpHandler(auth.HandleLambdaAuth(sessionStore)))
	})

	return r
}
