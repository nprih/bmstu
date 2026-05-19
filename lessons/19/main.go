package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var Status = map[string]string{"inStock": "на складе", "readyForIssue": "в ПВЗ"}
var LastIds = map[string]int{"user": 3, "order": 3}

var Users = []User{}
var Orders = []Order{}

type User struct {
	Id      int     `json:"id"`
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Age     int     `json:"age"`
	Balance int     `json:"balance"`
	Orders  []Order `json:"orders,omitempty"`
}

type Order struct {
	Id     int    `json:"id"`
	Length string `json:"length"`
	Width  string `json:"width"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Users)
}

func getOrdersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("orders"))
}

func createUsersHandler(w http.ResponseWriter, r *http.Request) {
	Users = append(Users, User{
		Id:      1,
		Name:    "",
		Email:   "",
		Age:     0,
		Balance: 100,
		Orders:  []Order{},
	})
}

func createOrdersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("orders"))
}

func addDefaultUsers() {
	Users = append(Users,
		User{
			Id:      1,
			Name:    "Adam",
			Email:   "adam@email.ru",
			Age:     46,
			Balance: 10000,
			Orders: append([]Order{}, Order{
				Id:     1,
				Length: "150",
				Width:  "30",
				Name:   "Shovel",
				Status: Status["readyForIssue"],
			}),
		},
		User{
			Id:      2,
			Name:    "Eva",
			Email:   "eva@email.ru",
			Age:     45,
			Balance: 15000,
			Orders: append([]Order{}, Order{
				Id:     2,
				Length: "200",
				Width:  "20",
				Name:   "Rake",
				Status: Status["readyForIssue"],
			}),
		},
		User{
			Id:      3,
			Name:    "Denis",
			Email:   "denis@email.ru",
			Age:     21,
			Balance: 5000,
			Orders: append([]Order{}, Order{
				Id:     3,
				Length: "3",
				Width:  "3",
				Name:   "Seeds",
				Status: Status["inStock`"],
			}),
		},
	)
}

func main() {
	addDefaultUsers()

	http.HandleFunc("GET /api/v1/users", getUsersHandler)
	http.HandleFunc("GET /api/v1/orders", getOrdersHandler)

	http.HandleFunc("POST /api/v1/users", createUsersHandler)
	http.HandleFunc("POST /api/v1/orders", createOrdersHandler)

	log.Println("Server starting...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
