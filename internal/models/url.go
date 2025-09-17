package models

import (
	"time"
)

type Url struct {
	ID        int    		`gorm:"primary_key;auto_increment" json:"id"`
	UserID    int    		`grom:"not null" json:"user_id"`
	ShortUrl  string 		`gorm:"not null;unique;size:255" json:"short_url"`
	LongUrl   string 		`gorm:"not null;size:255" json:"long_url"`
	CreatedAt time.Time 	`json:"created_at"`
	UpdatedAt time.Time 	`json:"updated_at"`
	DeletedAt time.Time 	`json:"deleted_at"`

	// Relasi balik: URL ini dimiliki oleh User
    User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
}