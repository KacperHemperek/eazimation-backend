package auth

import (
	"net/http"
	"time"
)

func SetSessionCookie(w http.ResponseWriter, sessionID string) {
	http.SetCookie(w, &http.Cookie{
		Domain:   getUrlWithoutProtocol(frontendUrl, isProd),
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
		Domain:   getUrlWithoutProtocol(frontendUrl, isProd),
		Name:     SessionCookieName,
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		Secure:   false,
		HttpOnly: true,
	})
}

func getUrlWithoutProtocol(url string, secure bool) string {
	if secure {
		return url[len("https://"):]
	}
	return url[len("http://"):]
}
