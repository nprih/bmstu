package handler

import "net/http"

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Index"))
}

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Task"))
}

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("CreateTask"))
}

func AscTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("AscTask"))
}
