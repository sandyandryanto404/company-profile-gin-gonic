package models

import (
	"time"
)

type UserNotification struct {
	Id          uint64    `json:"id" gorm:"primary_key"`
	UserId      uint64    `json:"user_id" gorm:"index;not null"`
	Subject     string    `json:"subject" gorm:"index;size:191;default:null;"`
	Description string    `json:"description" gorm:"index;size:255;default:null;"`
	Content     string    `json:"content"  gorm:"type:longtext;default:null;"`
	Status      uint8     `json:"status" gorm:"index;default:0"`
	CreatedAt   time.Time `gorm:"index;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"index;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (UserNotification) TableName() string {
	return "users_notifications"
}
