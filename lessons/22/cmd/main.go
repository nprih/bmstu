package main

import (
	"lesson22/internal/handler"
	"log"
	"net/http"
)

func main() {
	r := http.NewServeMux()

	r.HandleFunc("/", handler.IndexHandler)
	r.HandleFunc("/tasks", handler.TasksHandler)
	r.HandleFunc("/create-task", handler.CreateTaskHandler)
	r.HandleFunc("/ask-task", handler.AskTaskHandler)

	fs := http.FileServer(http.Dir("static"))
	r.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Server starting...")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Println(err)
	}
}
