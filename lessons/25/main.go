package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("INDEX PAGE, HELLO WITH RSA!!!)))"))
	})
	log.Println("Listening...")
	http.ListenAndServeTLS(":8085", "rootCA.pem", "rootCA.key", nil)
}
