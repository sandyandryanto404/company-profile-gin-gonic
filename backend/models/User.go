package models

import (
	"database/sql"
	"time"
)

type User struct {
	Id         uint64         `json:"id" gorm:"primary_key"`
	Image      string         `json:"image" gorm:"index;size:191;default:null;"`
	FirstName  string         `json:"first_name" gorm:"index;size:191;default:null;"`
	LastName   string         `json:"last_name" gorm:"index;size:191;default:null;"`
	Gender     string         `json:"gender" gorm:"index;size:2;default:null;"`
	Country    string         `json:"country" gorm:"index;size:191;default:null;"`
	Province   string         `json:"province" gorm:"index;size:191;default:null;"`
	Regency    string         `json:"regency" gorm:"index;size:191;default:null;"`
	PostalCode string         `json:"postal_code" gorm:"index;size:20;default:null;"`
	BirthPlace string         `json:"birth_place" gorm:"index;size:191;default:null;"`
	BirthDate  sql.NullTime   `gorm:"index;default:CURRENT_TIMESTAMP" json:"birth_date"`
	Address    sql.NullString `json:"address"  gorm:"type:text;default:null;"`
	AboutMe    sql.NullString `json:"about_me"  gorm:"type:text;default:null;"`
	Username   string         `json:"name" gorm:"index;size:191;not null"`
	Email      string         `json:"email" gorm:"index;size:191;not null"`
	Phone      string         `json:"phone" gorm:"index;size:191;default:null"`
	Password   string         `json:"password" gorm:"index;size:255;not null"`
	Status     uint8          `json:"status" gorm:"index;default:0"`
	CreatedAt  time.Time      `gorm:"index;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"index;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (User) TableName() string {
	return "users"
}
