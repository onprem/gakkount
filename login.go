package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func handleUserLogin(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": err.Error()})
		return
	}

	if !(strings.HasSuffix(input.Email, "iiitm.ac.in") && input.Password == "master") {
		c.JSON(http.StatusForbidden, gin.H{"status": "error", "error": "Invalid email or password"})
		return
	}

	token, err := createToken(input.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error": "Something went wrong"})
		log.Printf("login: create token: %s\n", err)
		return
	}

	c.SetCookie("token", token, 3600*48, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Successfully logged in"})
}
