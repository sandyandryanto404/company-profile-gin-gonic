package models

import (
	"time"
)

type Role struct {
	Id          uint64    `json:"id" gorm:"primary_key"`
	Code        string    `json:"code" gorm:"index;size:191;not null,unique"`
	Name        string    `json:"name" gorm:"index;size:191;not null,unique"`
	Description string    `json:"description"  gorm:"type:text;default:null;"`
	Status      uint8     `json:"status" gorm:"index;default:0"`
	CreatedAt   time.Time `gorm:"index;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"index;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (Role) TableName() string {
	return "roles"
}
