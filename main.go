package main

import (
	hd "auth/internals/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", hd.MainHandler)
	log.Println("Server (authAPI) started at http://localhost:8081")
	http.ListenAndServe(":8080", nil)
}
