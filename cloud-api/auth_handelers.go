package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var db *Database = NewDatabase()

func pingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func registerHandler(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	newUser := db.AddUser(&User{
		Email:    req.Email,
		Password: req.Password,
	})
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully", "user": newUser})
}

func usersHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"users": db.GetAllUsers()})
}

func loginHandler(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	user := db.GetUser(req.Email)
	if user != nil && user.Password == req.Password {
		c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": "dummy-token"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	}
}
