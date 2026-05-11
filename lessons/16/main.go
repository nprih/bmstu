package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	//if r.URL.Path != "/" {
	//	w.WriteHeader(http.StatusNotFound)
	//	return
	//}
	http.ServeFile(w, r, "form.html")
}
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About page"))
}
func contactsHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	action := params["action"]
	city := params["city"]

	switch {
	case action == "phone" && city == "msk":
		w.Write([]byte("Телефон в Москве: 8(495)0000000"))
		return
	case action == "phone" && city == "spb":
		w.Write([]byte("Телефон в Питерe: 8(812)0000000"))
		return
	case action == "address" && city == "msk":
		w.Write([]byte("Адрес в Москве: Беговая 19"))
		return
	case action == "address" && city == "spb":
		w.Write([]byte("Адрес в Питере: Большая монетная 3"))
		return
	}
}

func getHelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello my friend, i dont know ur email =("))
}

func postHelloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.PostFormValue("name")
	email := r.PostFormValue("email")
	fmt.Fprintf(w, "Hello %s, ur email: %s", name, email)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", indexHandler)
	router.HandleFunc("/about", aboutHandler)
	router.HandleFunc("/contacts/{action}/{city}", contactsHandler)
	router.HandleFunc("/hello", getHelloHandler).Methods("GET")
	router.HandleFunc("/hello", postHelloHandler).Methods("POST")

	log.Println("Server starting...")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
