package auth

import (
	"fmt"
	"github.com/gorilla/sessions"
	_ "github.com/joho/godotenv/autoload"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	"log/slog"
	"os"
)

var (
	maxAge             = 1000 * 60 * 60 * 24 * 30
	googleClientId     = os.Getenv("GOOGLE_CLIENT_ID")
	googleClientSecret = os.Getenv("GOOGLE_CLIENT_SECRET")
	sessionSecret      = os.Getenv("SESSION_SECRET")
	clientURL          = os.Getenv("FRONTEND_URL")
	apiPort            = os.Getenv("PORT")

	isProd = os.Getenv("APP_ENV") != "local"
)

func NewAuth() {
	store := sessions.NewCookieStore([]byte(sessionSecret))

	store.MaxAge(maxAge)
	store.Options.HttpOnly = true
	store.Options.Secure = false
	store.Options.Domain = clientURL

	gothic.Store = store
	googleCbURL := getGoogleCbURL()
	goth.UseProviders(google.New(googleClientId, googleClientSecret, googleCbURL))
}

func getGoogleCbURL() string {
	path := "api/v1/auth/google/callback"
	if isProd {
		slog.Info("using production google callback url")
		return fmt.Sprintf("https://ezm-api.kacperhemperek.com/%s", path)
	}
	slog.Info("using development google callback url")
	return fmt.Sprintf("http://localhost:%s/%s", apiPort, path)
}
