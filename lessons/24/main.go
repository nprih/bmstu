package main

import (
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
	c.String(http.StatusOK, "Status %s: ok", query)
}

func getUserById(c *gin.Context) {
	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	user := User{
		Id:    intId,
		Name:  "John Doe",
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
		"message": "User created",
		"user":    newUser,
	})
}

func main() {
	r := gin.Default()
	r.GET("/", indexHandler)
	r.GET("/user/: id", getUserById)
	r.POST("/user", createNewUser)
	r.GET("/status", statusHandler)

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
