package main

import (
	"log"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Write([]byte("Index page"))
}
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About page"))
}
func contactsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Contacts page"))
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/about", aboutHandler)
	mux.HandleFunc("/contacts", contactsHandler)

	log.Println("Server starting...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
