package models

import "time"

type ReferenceContent struct {
	Id          uint64    `json:"id" gorm:"primary_key"`
	Slug        string    `json:"slug" gorm:"index;size:255;not null"`
	Name        string    `json:"name" gorm:"index;size:255;not null"`
	Description string    `json:"description"  gorm:"type:text;not null"`
	Type        uint16    `json:"type" gorm:"index;default:0"`
	Status      uint8     `json:"status" gorm:"index;default:0"`
	CreatedAt   time.Time `gorm:"index;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"index;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (ReferenceContent) TableName() string {
	return "references_contents"
}
