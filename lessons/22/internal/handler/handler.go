package handler

import (
	"fmt"
	"html/template"
	"lesson22/internal/db"
	"log"
	"net/http"
	"strconv"
	"time"
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
	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("templates/create-task.html")

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
	} else if r.Method == http.MethodPost {
		id := r.FormValue("clientId")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		text := r.FormValue("text")
		newTask := db.Task{
			ClientId: idInt,
			Text:     text,
			Created:  time.Now().Format("02 Jan 2006"),
			Done:     0,
		}
		err = db.CreateNewTask(newTask)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/tasks", http.StatusFound)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func AscTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("AscTask"))
}
