package handlers

import (
	"auth/internals/tools"
	md "auth/models"
	"database/sql"
	"net/http"
)

func registerHandler(w http.ResponseWriter, user md.User, db *sql.DB) {
	err := user.CreateUser(db)
	if err != nil {
		http.Error(w, "User already exists", http.StatusConflict)
		return
	}
	tools.WriteResponse(w, "New user created", http.StatusCreated)
}
