package handler

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		"templates/css.html",
		"templates/head.html",
		"templates/header.html",
		"templates/main.html",
		"templates/js.html",
	)

	if err != nil {
		log.Println(err)
		fmt.Fprintln(w, err)
		return
	}

	err = tmpl.ExecuteTemplate(w, "main.html", "")
	if err != nil {
		log.Println(err)
		fmt.Fprintln(w, err)
		return
	}
}

func GetLoginHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		"templates/css.html",
		"templates/head.html",
		"templates/header.html",
		"templates/login.html",
		"templates/js.html",
	)
	if err != nil {
		log.Println(err)
		fmt.Fprintln(w, err)
		return
	}

	err = tmpl.ExecuteTemplate(w, "login.html", "")
	if err != nil {
		log.Println(err)
		fmt.Fprintln(w, err)
		return
	}
}

func PostLoginHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Can't read body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	json.NewEncoder(w).Encode(string(body))
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		"templates/css.html",
		"templates/head.html",
		"templates/header.html",
		"templates/register.html",
		"templates/js.html",
	)
	if err != nil {
		log.Println(err)
		fmt.Fprintln(w, err)
		return
	}

	err = tmpl.ExecuteTemplate(w, "register.html", "")
	if err != nil {
		log.Println(err)
		fmt.Fprintln(w, err)
		return
	}
}
