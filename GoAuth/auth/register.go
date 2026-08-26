package main

import (
	"net/http"

	"github.com/AimableUK/TheGopher/GoAuth/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

func registerHandler(c *gin.Context) {
	var req RegisterRequest

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		12,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to process your password",
		})
	}
	user := &models.User{
		Email:    req.Email,
		Password: string(hashedPassword),
		Role:     "user",
	}

	if err := models.UserStore.Create(user); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
	}

	user.Passwrd = ""
	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"user":    user,
	})

}
