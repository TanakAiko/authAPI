package main

import (
	hd "auth/internals/handlers"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	http.HandleFunc("/", hd.MainHandler)
	log.Println("Server (authAPI) started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
