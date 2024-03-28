package models

import (
	"time"
)

type LoginAttempt struct {
	Id        uint64    `json:"id" gorm:"primary_key"`
	IpAddress string    `json:"email" gorm:"index;size:45;not null"`
	Login     string    `json:"token" gorm:"index;size:255;not null"`
	CreatedAt time.Time `gorm:"index;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"index;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (LoginAttempt) TableName() string {
	return "login_attempts"
}
