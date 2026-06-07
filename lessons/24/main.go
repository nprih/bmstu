package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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

func main() {
	r := gin.Default()
	r.GET("/", indexHandler)
	r.GET("/user/:name", greetUserHandler)
	r.GET("/status", statusHandler)
	r.Run(":8080")
}
