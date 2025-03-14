package utils

import (
	"net/http"
	"os"
)

func CreateTokenCookie(token string) *http.Cookie {
	cookie := &http.Cookie{
		Name:     "Authorization",
		Value:    "Bearer " + token,
		Path:     "/",
		MaxAge:   3600 * 24 * 30,
		Secure:   os.Getenv("ENV") == "production",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}
	return cookie
}
