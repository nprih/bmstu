package main

import (
	"log"
	"net/http"
	"vkr/internal/handler"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.IndexHandler)
	mux.HandleFunc("/login", handler.LoginHandler)
	mux.HandleFunc("/register", handler.RegisterHandler)

	log.Println("Starting server...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Println(err)
	}
}
