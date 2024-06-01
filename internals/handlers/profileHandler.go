package handlers

import (
	"encoding/json"
	"net/http"
)

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(string)
	json.NewEncoder(w).Encode(map[string]string{"username": user})
}
