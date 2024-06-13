package main

import (
	hd "auth/internals/handlers"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	port := "8081"
	http.HandleFunc("/", hd.MainHandler)
	log.Printf("Server (authAPI) started at http://localhost:%v\n", port)
	http.ListenAndServe(":"+port, nil)
}
