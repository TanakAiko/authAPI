package handlers

import (
	"auth/internals/tools"
	md "auth/models"
	"database/sql"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func loginHandler(w http.ResponseWriter, user md.User, db *sql.DB) {
	hashedPassword, err := user.GetUser(db)
	if err != nil {
		http.Error(w, "Error: "+err.Error(), http.StatusUnauthorized)
		return
	}

	if !checkPasswordHash(user.Password, hashedPassword) {
		http.Error(w, "Error", http.StatusUnauthorized)
		return
	}

	var session md.Session
	err = session.CreateSession(user, db)
	if err != nil {
		http.Error(w, "Error: "+err.Error(), http.StatusConflict)
		return
	}

	user.SessionID = session.Id
	user.Password = ""
	http.SetCookie(w, &http.Cookie{
		Name:     "sessionID",
		Value:    session.Id,
		Path:     "/",
		Expires:  session.Expiration,
		HttpOnly: false,
	})
	user.SessionExpireTime = session.Expiration
	tools.WriteResponse(w, user, http.StatusOK)
}

func getUserData(w http.ResponseWriter, user md.User, db *sql.DB) {
	if err := user.GetUserFromSession(db); err != nil {
		http.Error(w, "Error: "+err.Error(), http.StatusUnauthorized)
		return
	}
	tools.WriteResponse(w, user, http.StatusOK)
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
