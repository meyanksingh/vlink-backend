package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/meyanksingh/vlink-backend/internal/app/models"
	"github.com/meyanksingh/vlink-backend/internal/app/repository"
	database "github.com/meyanksingh/vlink-backend/internal/db"
)

func Home(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
		return
	}

	//Populate Data in Golang
	var user models.User
	if err := database.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Error User Not Found "})
		return
	}

	// Send a welcome message with the user ID
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to the protected home page!",
		"email":   user.Email,
		"name":    user.FirstName + " " + user.LastName,
	})
}

func SendFriendRequest(c *gin.Context) {
	userID := c.MustGet("user_id").(uuid.UUID)
	var req struct {
		ReceiverID uuid.UUID `json:"receiver_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := repository.SendFriendRequest(userID, req.ReceiverID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Friend request sent"})
}

func AcceptFriendRequest(c *gin.Context) {
	userID := c.MustGet("user_id").(uuid.UUID)
	var req struct {
		RequestID uuid.UUID `json:"request_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	request, err := repository.GetFriendRequestByID(req.RequestID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if request.ReceiverID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to accept this request"})
		return
	}

	err = repository.AcceptFriendRequest(req.RequestID, userID, request.SenderID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Friend request accepted"})
}

func DeclineFriendRequest(c *gin.Context) {
	userID := c.MustGet("user_id").(uuid.UUID)
	var req struct {
		RequestID uuid.UUID `json:"request_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	request, err := repository.GetFriendRequestByID(req.RequestID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if request.ReceiverID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to decline this request"})
		return
	}

	err = repository.DeclineFriendRequest(req.RequestID, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Friend request declined"})
}

func RemoveFriend(c *gin.Context) {
	userID := c.MustGet("user_id").(uuid.UUID)
	var req struct {
		FriendID uuid.UUID `json:"friend_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := repository.RemoveFriend(userID, req.FriendID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Friend removed"})
}

func ListFriends(c *gin.Context) {
	userID := c.MustGet("user_id").(uuid.UUID)
	friends, err := repository.ListFriends(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve friends"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"friends": friends})
}

func ListFriendRequests(c *gin.Context) {
	userID := c.MustGet("user_id").(uuid.UUID)
	requests, err := repository.ListFriendRequests(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve friend requests"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"friend_requests": requests})
}
