package models

import (
	"time"
)

type PasswordReset struct {
	Id        uint64    `json:"id" gorm:"primary_key"`
	Email     string    `json:"email" gorm:"index;size:191;not null"`
	Token     string    `json:"token" gorm:"index;size:255;not null"`
	Status    uint8     `json:"status" gorm:"index;default:0"`
	CreatedAt time.Time `gorm:"index;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"index;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (PasswordReset) TableName() string {
	return "password_resets"
}
