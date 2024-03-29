package models

import "time"

type Contact struct {
	Id        uint64    `json:"id" gorm:"primary_key"`
	Name      string    `json:"name" gorm:"index;size:255;not null"`
	Email     string    `json:"email" gorm:"index;size:255;not null"`
	Subject   string    `json:"subject" gorm:"index;size:255;not null"`
	Message   string    `json:"content"  gorm:"type:text;not null"`
	Status    uint8     `json:"status" gorm:"index;default:0"`
	CreatedAt time.Time `gorm:"index;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"index;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (Contact) TableName() string {
	return "contacts"
}
