package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Signup Successfull"})

}

func Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Login Successfull"})

}
