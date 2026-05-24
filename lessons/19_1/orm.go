package main

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type user struct {
	gorm.Model
	Firstname string
	Lastname  string
	Email     string
	Age       int
}

func main() {
	pgConfig := "host=localhost user=postgres password=changeme dbname=postgres port=5432"

	db, err := gorm.Open(postgres.Open(pgConfig), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return
	}

}
