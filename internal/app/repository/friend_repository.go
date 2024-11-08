package repository

import (
	"errors"

	"github.com/google/uuid"
	"github.com/meyanksingh/vlink-backend/internal/app/models"
	db "github.com/meyanksingh/vlink-backend/internal/db"
)

func SendFriendRequest(senderID, receiverID uuid.UUID) error {
	if senderID == receiverID {
		return errors.New("cannot send friend request to yourself")
	}
	var user models.User
	if err := db.DB.Where("email = ?", receiverID).First(&user).Error; err != nil {
		return errors.New("friend not found")
	}
	var count int64
	err := db.DB.Model(&models.Friend{}).
		Where("(user_id = ? AND friend_id = ?) OR (user_id = ? AND friend_id = ?)", senderID, receiverID, receiverID, senderID).
		Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("you are already friends with this user")
	}

	err = db.DB.Model(&models.FriendRequest{}).
		Where("sender_id = ? AND receiver_id = ? AND status = ?", senderID, receiverID, "pending").
		Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("friend request already exists")
	}

	friendRequest := models.FriendRequest{
		SenderID:   senderID,
		ReceiverID: receiverID,
		Status:     "pending",
	}
	return db.DB.Create(&friendRequest).Error
}

func AcceptFriendRequest(requestID, userID, friendID uuid.UUID) error {
	var request models.FriendRequest
	err := db.DB.Where("id = ? AND receiver_id = ? AND status = ?", requestID, userID, "pending").First(&request).Error
	if err != nil {
		return errors.New("friend request not found or already processed")
	}

	err = db.DB.Model(&models.FriendRequest{}).
		Where("id = ?", requestID).
		Update("status", "accepted").Error
	if err != nil {
		return err
	}

	friend := models.Friend{
		UserID:   userID,
		FriendID: friendID,
	}
	return db.DB.Create(&friend).Error
}

func DeclineFriendRequest(requestID, userID uuid.UUID) error {
	var request models.FriendRequest
	err := db.DB.Where("id = ? AND receiver_id = ? AND status = ?", requestID, userID, "pending").First(&request).Error
	if err != nil {
		return errors.New("friend request not found or already accepted/rejected")
	}

	return db.DB.Model(&models.FriendRequest{}).
		Where("id = ?", requestID).
		Update("status", "declined").Error
}

func RemoveFriend(userID, friendID uuid.UUID) error {
	var count int64
	err := db.DB.Model(&models.Friend{}).
		Where("(user_id = ? AND friend_id = ?) OR (user_id = ? AND friend_id = ?)", userID, friendID, friendID, userID).
		Count(&count).Error
	if err != nil {
		return err
	}

	if count == 0 {
		return errors.New("you are not friends with this user")
	}

	return db.DB.Where("(user_id = ? AND friend_id = ?) OR (user_id = ? AND friend_id = ?)", userID, friendID, friendID, userID).
		Delete(&models.Friend{}).Error
}

func ListFriends(userID uuid.UUID) ([]models.Friend, error) {
	var friends []models.Friend
	err := db.DB.Where("user_id = ? OR friend_id = ?", userID, userID).Find(&friends).Error
	return friends, err
}

func ListFriendRequests(userID uuid.UUID) ([]models.FriendRequest, error) {
	var requests []models.FriendRequest
	err := db.DB.Where("receiver_id = ? AND status = ?", userID, "pending").Find(&requests).Error
	return requests, err
}

func GetFriendRequestByID(requestID uuid.UUID) (models.FriendRequest, error) {
	var request models.FriendRequest
	err := db.DB.Where("id = ?", requestID).First(&request).Error
	if err != nil {
		return request, err
	}
	return request, nil
}
