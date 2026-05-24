package main

import (
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

	/*
		INSERT
	*/
	//dmitry := User{
	//	Firstname: "Ivan",
	//	Lastname:  "Kukushkin",
	//	Email:     "Ivan@mail.ru",
	//	Age:       30,
	//}
	//res := db.Create(&dmitry)
	//if res.Error != nil {
	//	log.Println(res.Error)
	//	return
	//}
	//fmt.Println("User", dmitry.Firstname, "added to db")

	/*
		SELECT + WHERE
	*/
	var users []User
	//res := db.Where("firstname = ?", "Dmitry").Find(&users)
	res := db.Find(&users)
	if res.Error != nil {
		log.Println(res.Error)
		return
	}
	for _, user := range users {
		log.Printf("Name: %s, Surname: %s, Age: %d", user.Firstname, user.Lastname, user.Age)
	}

	/*
		UPDATE
	*/
	//var user User
	//res := db.First(&user, 1)
	//if res.Error != nil {
	//	fmt.Println(res.Error)
	//	return
	//}
	//user.Age = 20
	//user.Firstname = "Andrew"
	//res = db.Save(&user)
	//if res.Error != nil {
	//	fmt.Println(res.Error)
	//	return
	//}

	/*
		DELETE
	*/
	//var user User
	//res := db.First(&user, 1)
	//if res.Error != nil {
	//	fmt.Println(res.Error)
	//	return
	//}
	//res = db.Delete(&user)
	//if res.Error != nil {
	//	fmt.Println(res.Error)
	//} else if res.RowsAffected == 0 {
	//	fmt.Println("No such user")
	//} else {
	//	fmt.Println("Delete Success")
	//}
}
