package main

import (
	"log"
	"net/http"
	"vkr/internal/handler"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.IndexHandler)
	mux.HandleFunc("GET /login", handler.GetLoginHandler)
	mux.HandleFunc("POST /login", handler.PostLoginHandler)
	mux.HandleFunc("/register", handler.RegisterHandler)

	log.Println("Starting server...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Println(err)
	}
}
