package auth

import (
	"context"
	"eazimation-backend/internal/api"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func NewAddProviderToContext() AddProviderToContextMiddleware {
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

type AddProviderToContextMiddleware = func(h api.HandlerFunc) api.HandlerFunc
