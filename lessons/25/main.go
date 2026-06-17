package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"log"
	"net/http"
	"time"
)

var secretKey = []byte("myexamplesimplekey")
var database = map[string]string{
	"Misha": "password123",
	"Petya": "123456",
	"Sasha": "654321",
}

type LoginRequest struct {
	Username string `json:"name"`
	Password string `json:"password"`
}

func login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	log.Println("New login try:", req.Username)
	correctPassword, exists := database[req.Username]
	if !exists || correctPassword != req.Password {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid username or password",
		})
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name":    req.Username,
		"created": time.Now().Unix(),
		"expired": time.Now().Add(24 * time.Hour).Unix(),
	})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Cannot create JWT, sorry =(",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"token":   tokenString,
	})
}

func secret(c *gin.Context) {
	var tokenString string
	authHeader := c.GetHeader("Authorization")
	if authHeader != "" {
		if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
			tokenString = authHeader[7:]
		}
	}
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "use Authorization:Bearer ur_token",
		})
		return
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Invalid sign method")
		}
		return secretKey, nil
	})
	if err != nil {
		log.Println("validation error", err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":       "Token is not valid",
			"description": err.Error(),
		})
		return
	}
	if climes, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		log.Println("Token is VALID!")
		c.JSON(http.StatusOK, gin.H{
			"message": "This is very secret page =)",
			"name":    climes["name"],
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
	}
}

func users(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello! This is public page!",
	})
}

func main() {
	router := gin.Default()
	router.POST("/login", login)
	router.GET("/secret", secret)
	router.GET("/users", users)
	if err := router.Run(":8085"); err != nil {
		log.Println(err)
		return
	}
}
