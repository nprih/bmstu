package main

import (
	"log"
	"net/http"
)

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "ok"}`))
}
func main() {
	http.HandleFunc("/status", StatusHandler)
	log.Println("Server starting...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Println(err)
		return
	}
}
