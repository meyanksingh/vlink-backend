package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/meyanksingh/vlink-backend/internal/app/repository"
	"github.com/meyanksingh/vlink-backend/pkg/utils"
)

func Register(c *gin.Context) {
	var requestBody struct {
		FirstName string `json:"first_name" binding:"required"`
		LastName  string `json:"last_name" binding:"required"`
		Email     string `json:"email" binding:"required,email"`
		Password  string `json:"password" binding:"required,min=8"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	emailExists, err := repository.CheckEmailExists(requestBody.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking email"})
		return
	}
	if emailExists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already in use"})
		return
	}

	user, err := repository.CreateUser(requestBody.FirstName, requestBody.LastName, requestBody.Email, requestBody.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Signup successful",
		"email":   user.Email,
	})
}

func Login(c *gin.Context) {
	var requestBody struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	user, err := repository.AuthenticateUser(requestBody.Email, requestBody.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.ID.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
		"user":    user,
	})
}
