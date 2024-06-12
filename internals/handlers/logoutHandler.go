package handlers

import (
	"auth/internals/tools"
	md "auth/models"
	"database/sql"
	"net/http"
)

func logoutHandler(w http.ResponseWriter, user md.User, db *sql.DB) {
	session := md.Session{}
	session.Id = user.SessionID
	session.DeleteSession(db)
	tools.WriteResponse(w, "The session is deleted", http.StatusOK)
}
