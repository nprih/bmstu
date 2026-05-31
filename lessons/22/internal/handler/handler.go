package handler

import (
	"fmt"
	"html/template"
	"lesson22/internal/db"
	"log"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/home.html")

	if err != nil {
		log.Println(err)
		fmt.Fprintln(w, err)
		return
	}
	clients := db.SelectAllClients()
	if err = tmpl.Execute(w, clients); err != nil {
		log.Println(err)
		fmt.Fprintln(w, err)
		return
	}
}

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/tasks.html")

	if err != nil {
		log.Println(err)
		fmt.Fprintln(w, err)
		return
	}
	tasks := db.SelectAllTask()
	if err = tmpl.Execute(w, tasks); err != nil {
		log.Println(err)
		fmt.Fprintln(w, err)
		return
	}
}

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/create-task.html")

	if err != nil {
		log.Println(err)
		fmt.Fprintln(w, err)
		return
	}

	if err = tmpl.Execute(w, ""); err != nil {
		log.Println(err)
		fmt.Fprintln(w, err)
		return
	}
}

func AscTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("AscTask"))
}
