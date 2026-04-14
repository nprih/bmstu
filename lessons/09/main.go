package main

import (
	"encoding/json"
	"fmt"
)

type Car struct {
	Number string `json:"car_number"`
	Color  string `json:"car_color"`
	Engine string `json:"car_engine"`
}

func main() {
	myCar := Car{
		Number: "а001аа777",
		Color:  "red",
		Engine: "B123A56",
	}

	jsonData, err := json.Marshal(myCar)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jsonData))
}
