package handler

import (
	"encoding/json"
	"fmt"
	"lesson22/internal/db"
	"lesson22/internal/service"
	"log"
	"net/http"
	"strconv"
	"text/template"
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

func TasksHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/tasks.html")
	if err != nil {
		log.Println(err)
		fmt.Fprintln(w, err)
		return
	}
	tasks := db.SelectAllTasks()
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

func AskTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		queryParams := r.URL.Query()
		clientId := queryParams.Get("id")
		IntClientId, err := strconv.Atoi(clientId)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		isExists := service.CheckClientById(IntClientId)
		if !isExists {
			err = db.InsertNewClient(IntClientId)
			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
			}
		}

		tasks := db.SelectAllTasks()
		type Task struct {
			Id   int    //0
			Text string //""
		}
		taskForClient := Task{}
		for _, v := range tasks {
			if v.ClientId == IntClientId && v.Done == 0 {
				taskForClient.Id = v.Id
				taskForClient.Text = v.Text
				break
			}
		}
		w.Header().Set("Content-Type", "application/json")
		res, err := json.Marshal(taskForClient)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(res)
	} else if r.Method == http.MethodPost {

		var clientTaskAnswer db.TaskAnswer
		err := json.NewDecoder(r.Body).Decode(&clientTaskAnswer)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = db.UpdateTask(clientTaskAnswer)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
