package middlewares

import (
	"context"
	"net/http"

	"github.com/ckwcfm/learn-go/rss/constants"
	"github.com/ckwcfm/learn-go/rss/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func IsUser(next http.Handler) http.Handler {
	return Authorization(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.Context().Value(constants.UserIDKey).(string)
		userID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		user, err := services.GetUserByID(userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), constants.UserKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	}))
}
