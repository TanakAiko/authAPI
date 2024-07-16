package handlers

import (
	dbManager "auth/internals/dbManager"
	md "auth/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	db, err := dbManager.InitDB()
	if err != nil {
		log.Println("db not opening !", err)
		http.Error(w, "database can't be opened", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var req md.Request
	json.NewDecoder(r.Body).Decode(&req)

	fmt.Println("req.Action : ", req.Action)

	switch req.Action {
	case "register":
		registerHandler(w, req.Body, db)
	case "login":
		loginHandler(w, req.Body, db)
	case "logout":
		logoutHandler(w, req.Body, db)
	case "authorized":
		authorized(w, req.Body, db)
	case "getUserData":
		getUserData(w, req.Body, db)
	case "getAllUser":
		getAllUser(w, db)

	default:
		http.Error(w, "Unknown action", http.StatusBadRequest)
		return
	}
}
