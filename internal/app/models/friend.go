package models

import (
	"time"

	"github.com/google/uuid"
)

type Friend struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	FriendID  uuid.UUID `json:"friend_id" gorm:"type:uuid;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}

type FriendRequest struct {
	ID         uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	SenderID   uuid.UUID `json:"sender_id" gorm:"type:uuid;not null"`
	ReceiverID uuid.UUID `json:"receiver_id" gorm:"type:uuid;not null"`
	Status     string    `json:"status" gorm:"type:varchar(20);not null"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
}
