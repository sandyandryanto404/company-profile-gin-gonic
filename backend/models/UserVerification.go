package models

import (
	"time"
)

type UserVerification struct {
	Id        uint64    `json:"id" gorm:"primary_key"`
	UserId    uint64    `json:"user_id" gorm:"index;not null"`
	Token     string    `json:"token" gorm:"index;size:255;not null"`
	Status    uint8     `json:"status" gorm:"index;default:0"`
	ExpiredAt time.Time `gorm:"index;default:CURRENT_TIMESTAMP" json:"expired_at"`
	CreatedAt time.Time `gorm:"index;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"index;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (UserVerification) TableName() string {
	return "users_verifications"
}
