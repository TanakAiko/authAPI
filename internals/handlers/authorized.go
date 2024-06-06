package handlers

import (
	md "auth/models"
	"database/sql"
	"net/http"
)

func authorized(w http.ResponseWriter, user md.User, db *sql.DB) {
	var session md.Session
	session.Id = user.SessionID
	if err := session.TestSessionUser(db); err != nil {
		http.Error(w, "Forbidden : "+err.Error(), http.StatusForbidden)
	}

	w.WriteHeader(http.StatusAccepted)
}
