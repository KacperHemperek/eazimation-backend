package auth

import (
	"net/http"
	"time"
)

func SetSessionCookie(w http.ResponseWriter, sessionID string) {
	http.SetCookie(w, &http.Cookie{
		Name:     SessionCookieName,
		Value:    sessionID,
		Path:     "/",
		MaxAge:   int(time.Hour * 24 * 7),
		Secure:   false,
		HttpOnly: true,
	})
}

func RemoveSessionCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     SessionCookieName,
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		Secure:   false,
		HttpOnly: true,
	})
}
