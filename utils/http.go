package utils

import (
	"encoding/json"
	"net/http"
)

func GetQueryWithDefault[T comparable](r *http.Request, key string, defaultValue T) T {
	query := r.URL.Query().Get(key)
	if query == "" {
		return defaultValue
	}
	var result T
	if err := json.Unmarshal([]byte(query), &result); err != nil {
		return defaultValue
	}
	return result
}

// func GetQueryWithDefault(r *http.Request, key string, defaultValue string) string {
// 	query := r.URL.Query().Get(key)
// 	if query == "" {
// 		return defaultValue
// 	}
// 	return query
// }
