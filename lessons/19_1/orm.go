package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Firstname string `gorm:"not null"`
	Lastname  string `gorm:"not null"`
	Email     string `gorm:"uniqueIndex"`
	Age       int    `gorm:"size:3"`
}

func main() {
	pgConfig := "host=localhost user=postgres password=changeme dbname=postgres port=5432"

	db, err := gorm.Open(postgres.Open(pgConfig), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return
	}
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Println(err)
		return
	}

	dmitry := User{
		Firstname: "Dmitry",
		Lastname:  "Verevkin",
		Email:     "dmitriy@mail.ru",
		Age:       30,
	}
	res := db.Create(&dmitry)
	if res.Error != nil {
		log.Println(res.Error)
		return
	}
	fmt.Println("User", dmitry.Firstname, "added to db")
}
