package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id    int
	Name  string
	Email string
}

func indexHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello from gin")
}

func statusHandler(c *gin.Context) {
	query := c.DefaultQuery("id", "no service")
	c.String(http.StatusOK, "Status for %s: ok", query)
}

func getUserById(c *gin.Context) {
	id := c.Param("id")
	IntId, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	user := User{
		Id:    IntId,
		Name:  "Vitaly",
		Email: "Vitaly@mail.ru",
	}
	c.JSON(http.StatusOK, user)
}

func createNewUser(c *gin.Context) {
	var newUser User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User was created",
		"user":    newUser,
	})
}

func AuthMiddeware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("New Auth Request:", c.Request.Method, c.Request.URL.Path)
		c.Next()
	}
}

func AdminMiddeware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("ONLY FOR ADMIN!!")
		c.Next()
	}
}

func main() {
	r := gin.Default()
	r.GET("/", indexHandler)

	r.Use(AuthMiddeware())
	r.GET("/user/:id", getUserById)
	r.POST("/user", createNewUser)

	r.Use(AdminMiddeware())
	r.GET("/status", statusHandler)

	if err := r.Run(":8080"); err != nil {
		log.Println(err)
		return
	}
}
