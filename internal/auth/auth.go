package auth

import (
	"fmt"
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
	apiPort            = os.Getenv("PORT")
	isProd             = os.Getenv("APP_ENV") == "production"
)

func NewAuth() {
	store := sessions.NewCookieStore([]byte(sessionSecret))

	store.MaxAge(maxAge)
	store.Options.HttpOnly = true
	store.Options.Secure = false
	store.Options.Domain = getStoreDomain()

	gothic.Store = store
	googleCbURL := getGoogleCbURL()
	goth.UseProviders(google.New(googleClientId, googleClientSecret, googleCbURL))
}

func getGoogleCbURL() string {
	path := "api/v1/auth/google/callback"
	if isProd {
		return fmt.Sprintf("https://ezm-api.kacperhemperek.com/%s", path)
	}
	return fmt.Sprintf("http://localhost:%s/%s", apiPort, path)
}

func getStoreDomain() string {
	if isProd {
		return "ezm-api.kacperhemperek.com"
	}
	return "localhost"
}
