package auth

import (
	"github.com/gorilla/sessions"
	_ "github.com/joho/godotenv/autoload"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	"os"
)

var (
	maxAge             = 1000 * 60 * 60 * 24 * 30
	googleClientId     = os.Getenv("GOOGLE_CLIENT_ID")
	googleClientSecret = os.Getenv("GOOGLE_CLIENT_SECRET")
	sessionSecret      = os.Getenv("SESSION_SECRET")

	googleCallbackUrl = "http://localhost:8080/api/v1/auth/google/callback"
	clientURL         = "http://localhost:5173"
)

func NewAuth() {
	store := sessions.NewCookieStore([]byte(sessionSecret))

	store.MaxAge(maxAge)
	store.Options.HttpOnly = true
	store.Options.Secure = false
	store.Options.Domain = clientURL

	gothic.Store = store

	goth.UseProviders(google.New(googleClientId, googleClientSecret, googleCallbackUrl))
}
