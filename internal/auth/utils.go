package auth

import (
	"net/http"
	"time"
)

func SetSessionCookie(w http.ResponseWriter, sessionID string) {
	http.SetCookie(w, &http.Cookie{
		Name:     SessionCookieName,
		Value:    sessionID,
		Domain:   getCookieDomain(),
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
		Domain:   getCookieDomain(),
		Path:     "/",
		MaxAge:   -1,
		Secure:   isProd,
		HttpOnly: true,
	})
}

func getCookieDomain() string {
	if isProd {
		return "kacperhemperek.com"
	}
	return "localhost"
}
