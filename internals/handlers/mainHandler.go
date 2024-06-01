package handlers

import (
	md "auth/models"
	"encoding/json"
	"net/http"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	var req md.Request
	json.NewDecoder(r.Body).Decode(&req)

	switch req.Action {
	case "register":
		registerHandler(w, req)
	case "login":
		loginHandler(w, req)
	case "profile":
		authMiddleware(profileHandler)(w, r)
	default:
		http.Error(w, "Unknown action", http.StatusBadRequest)
	}
}
