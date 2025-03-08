package middlewares

import (
	"context"
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/ckwcfm/learn-go/rss/constants"
	"github.com/ckwcfm/learn-go/rss/services"
)

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("Authorization")
		log.Println(cookie)
		if err != nil {
			Unauthorized(next).ServeHTTP(w, r)
			return
		}
		token := strings.TrimPrefix(cookie.Value, "Bearer ")
		log.Println(token)
		if token == "" {
			Unauthorized(next).ServeHTTP(w, r)
			return
		}
		userID, err := services.ValidateToken(token)
		if err != nil {
			Unauthorized(next).ServeHTTP(w, r)
			return
		}
		ctx := context.WithValue(r.Context(), constants.UserIDKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func Unauthorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tmpls := []string{
			"views/layouts/main.html",
			"views/pages/unauthorized.html",
		}
		tmpl := template.Must(template.ParseFiles(tmpls...))
		tmpl.Execute(w, nil)
	})
}
