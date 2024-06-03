package handlers

import (
	dbManager "auth/internals/dbManager"
	md "auth/models"
	"database/sql"
	"net/http"
)

func registerHandler(w http.ResponseWriter, user md.User, db *sql.DB) {
	err := dbManager.CreateUser(user, db)
	if err != nil {
		http.Error(w, "User already exists", http.StatusConflict)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
