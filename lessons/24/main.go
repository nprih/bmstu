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

func greetUserHandler(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "Hello, %s", name)
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

func main() {
	r := gin.Default()
	r.GET("/", indexHandler)
	//r.GET("/user/:name", greetUserHandler)
	r.GET("/user/: id", getUserById)
	r.GET("/status", statusHandler)

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
