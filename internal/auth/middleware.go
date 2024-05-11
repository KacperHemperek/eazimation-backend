package auth

import (
	"context"
	"eazimation-backend/internal/api"
	"errors"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func NewAddProviderToContext() Middleware {
	return func(h api.HandlerFunc) api.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) error {
			provider := chi.URLParam(r, "provider")
			if provider == "" {
				return &api.Error{
					Message: "Invalid route, provider not found in URL",
					Code:    http.StatusInternalServerError,
				}
			}

			r = r.WithContext(context.WithValue(r.Context(), "provider", provider))
			return h(w, r)
		}
	}
}

func NewAuthMiddleware(sessionStore SessionStore) Middleware {
	return func(h api.HandlerFunc) api.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) error {
			sessionCookie, err := r.Cookie(SessionCookieName)
			var sessionID string
			if err != nil {
				sessionID = r.Header.Get("SessionId")

				if sessionID == "" {
					return NewUnauthorizedApiError(errors.New("authentication header not found"))
				}
			} else {
				sessionID = sessionCookie.Value
			}
			session, err := sessionStore.GetSession(sessionID)

			if err != nil {
				return NewUnauthorizedApiError(err)
			}
			r = r.WithContext(context.WithValue(r.Context(), "session", session))
			return h(w, r)
		}
	}
}

type Middleware = func(h api.HandlerFunc) api.HandlerFunc
