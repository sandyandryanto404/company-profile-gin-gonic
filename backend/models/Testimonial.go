package models

import "time"

type Testimonial struct {
	Id         uint64    `json:"id" gorm:"primary_key"`
	CustomerId uint64    `json:"customer_id" gorm:"index;not null"`
	Image      string    `json:"image" gorm:"index;size:191;default:null;"`
	Name       string    `json:"name" gorm:"index;size:255;not null"`
	Position   string    `json:"position" gorm:"index;size:255;not null"`
	Quote      string    `json:"quote" gorm:"type:text;default:null;"`
	Sort       uint16    `json:"sort" gorm:"index;default:0"`
	Status     uint8     `json:"status" gorm:"index;default:0"`
	CreatedAt  time.Time `gorm:"index;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"index;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (Testimonial) TableName() string {
	return "testimonials"
}
