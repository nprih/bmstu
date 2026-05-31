package handler

import (
	"fmt"
	"html/template"
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

func LoginHandler(w http.ResponseWriter, r *http.Request) {
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
