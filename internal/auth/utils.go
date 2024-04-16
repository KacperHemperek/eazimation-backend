package auth

import (
	"net/http"
	"os"
	"time"
)

var (
	feDomain = os.Getenv("FRONTEND_DOMAIN")
)

func SetSessionCookie(w http.ResponseWriter, sessionID string) {
	http.SetCookie(w, &http.Cookie{
		Name:     SessionCookieName,
		Value:    sessionID,
		Domain:   feDomain,
		Path:     "/",
		MaxAge:   int(time.Hour * 24 * 7),
		Secure:   isProd,
		HttpOnly: true,
	})
}

func RemoveSessionCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     SessionCookieName,
		Value:    "",
		Domain:   feDomain,
		Path:     "/",
		MaxAge:   -1,
		Secure:   isProd,
		HttpOnly: true,
	})
}
