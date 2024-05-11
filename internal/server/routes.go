package server

import (
	"eazimation-backend/internal/api"
	"eazimation-backend/internal/auth"
	"eazimation-backend/internal/handlers/auth"
	"eazimation-backend/internal/handlers/health"
	"eazimation-backend/internal/handlers/video"
	"eazimation-backend/internal/services/user"
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
	addProviderToCtx auth.Middleware,
	authMiddleware auth.Middleware,
	serverAuthMiddleware auth.Middleware,
	sessionStore auth.SessionStore,
	userService services.UserService,
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
		r.Get("/health", api.HttpHandler(healthhandlers.GetHealthHandler()))

		r.Get("/auth/{provider}/callback", api.HttpHandler(addProviderToCtx(authhandlers.HandleAuthCallback(sessionStore, userService))))
		r.Get("/auth/{provider}", api.HttpHandler(addProviderToCtx(authhandlers.HandleAuth(sessionStore, userService))))
		r.Post("/auth/logout", api.HttpHandler(authhandlers.HandleLogout(sessionStore)))
		r.Get("/auth/user", api.HttpHandler(authMiddleware(authhandlers.HandleGetUser())))
		r.Get("/auth/lambda", api.HttpHandler(authhandlers.HandleLambdaAuth(sessionStore)))

		r.Post("/videos", api.HttpHandler(serverAuthMiddleware(videohandlers.HandleCreateVideo())))
	})

	return r
}
