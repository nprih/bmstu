package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type User struct {
	Name string
	City string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	//user := User{Name: "Peter", City: "Mahachcala"}
	time := 10
	tmpl, err := template.ParseFiles("templates/index.html")

	if err != nil {
		log.Println(err)
		fmt.Fprintln(w, err)
		return
	}

	if err = tmpl.Execute(w, time); err != nil {
		log.Println(err)
		fmt.Fprintln(w, err)
		return
	}
}
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is about page")
}
func contactsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is contacts page")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/contacts", contactsHandler)
	log.Println("Server starting...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
