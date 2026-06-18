package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	version := os.Getenv("VERSION")
	if version == "" {
		version = "1.0.0"
	}
	fmt.Fprintf(w, "Heloo from docker (version: %s)\n", version)
}

func main() {
	http.HandleFunc("/status", IndexHandler)
	log.Println("Server starting...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Println(err)
		return
	}
}
