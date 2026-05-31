package main

import (
	"lesson22/internal/handler"
	"log"
	"net/http"
)

func main() {
	r := http.NewServeMux()
	r.HandleFunc("/", handler.IndexHandler)
	r.HandleFunc("/tasks", handler.TaskHandler)
	r.HandleFunc("/create-task", handler.CreateTaskHandler)
	r.HandleFunc("/ask-task", handler.AscTaskHandler)

	log.Println("Starting server...")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Println(err)
	}
}
