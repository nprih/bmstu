package main

import (
	"database/sql"
	"fmt"
)
import _ "modernc.org/sqlite"
import _ "github.com/lib/pq"

type Phone struct {
	Id     int
	Vendor string
	Model  string
	Price  int
}

func main() {
	pgConfig := "user=postgres password=changeme dbname=postgres sslmode=disable port=5432"
	db, err := sql.Open("postgres", pgConfig)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("ping", err)
	}

	_, err = db.Exec("INSERT INTO phones (vendor, model, price) VALUES ($1, $2, $3)", "Apple", "iPhone 12", 7000)
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
