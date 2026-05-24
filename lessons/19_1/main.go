package main

import (
	"database/sql"
	"fmt"
)
import _ "modernc.org/sqlite"

type Phone struct {
	Id     int
	Vendor string
	Model  string
	Price  int
}

func main() {
	db, err := sql.Open("sqlite", "phones")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO phones (vendor, model, price) VALUES (?, ?, ?)", "Apple", "iPhone 12", 7000)
	if err != nil {
		fmt.Println(err)
		return
	}

	rows, err := db.Query("SELECT * FROM phones")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()
	phones := []Phone{}
	for rows.Next() {
		ph := Phone{}
		err := rows.Scan(&ph.Id, &ph.Vendor, &ph.Model, &ph.Price)
		if err != nil {
			fmt.Println(err)
			continue
		}
		phones = append(phones, ph)
	}
	fmt.Println(phones)
}
