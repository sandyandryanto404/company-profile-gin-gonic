package models

import (
	"database/sql"
	"time"
)

type Portfolio struct {
	Id          uint64       `json:"id" gorm:"primary_key"`
	CustomerId  uint64       `json:"user_id" gorm:"index;not null"`
	ReferenceId uint64       `json:"reference_id" gorm:"index;not null"`
	Title       string       `json:"title" gorm:"index;size:255;not null"`
	Description string       `json:"description"  gorm:"type:text;not null"`
	ProjectDate sql.NullTime `gorm:"index;default:CURRENT_TIMESTAMP" json:"project_date"`
	ProjectUrl  string       `json:"project_url" gorm:"type:text;default:null;"`
	Sort        uint16       `json:"sort" gorm:"index;default:0"`
	Status      uint8        `json:"status" gorm:"index;default:0"`
	CreatedAt   time.Time    `gorm:"index;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time    `gorm:"index;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (Portfolio) TableName() string {
	return "portfolios"
}
