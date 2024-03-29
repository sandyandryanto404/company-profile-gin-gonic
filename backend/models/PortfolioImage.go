package models

import "time"

type PortfolioImage struct {
	Id          uint64    `json:"id" gorm:"primary_key"`
	PortfolioId uint64    `json:"portfolio_id" gorm:"index;not null"`
	Image       string    `json:"image" gorm:"index;size:255;not null"`
	Status      uint8     `json:"status" gorm:"index;default:0"`
	CreatedAt   time.Time `gorm:"index;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"index;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (PortfolioImage) TableName() string {
	return "portfolios_images"
}
