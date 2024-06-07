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
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	if !checkPasswordHash(user.Password, hashedPassword) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	var session md.Session
	err = session.CreateSession(user, db)
	if err != nil {
		http.Error(w, "Error creating session", http.StatusInternalServerError)
		return
	}

	user.SessionID = session.Id
	user.Password = ""
	http.SetCookie(w, &http.Cookie{
		Name:    "sessionID",
		Value:   session.Id,
		Path:    "/",
		Expires: session.Expiration,
	})

	tools.WriteResponse(w, "", http.StatusOK)
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
