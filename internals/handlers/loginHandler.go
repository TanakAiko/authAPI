package handlers

import (
	md "auth/models"
	"database/sql"
	"encoding/json"
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

	http.SetCookie(w, &http.Cookie{
		Name:    "sessionID",
		Value:   session.Id,
		Path:    "/",
		Expires: session.Expiration,
	})

	w.WriteHeader(http.StatusOK)
	data, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Error : Marshal data to send", http.StatusInternalServerError)
		return
	}
	_, err = w.Write(data)
	if err != nil {
		http.Error(w, "Error : Writing the data to the response", http.StatusInternalServerError)
		return
	}
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
