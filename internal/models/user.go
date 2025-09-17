package models

import "time"

type UserInMemory struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type User struct {
    ID        uint      `gorm:"primaryKey"`
    Username  string    `gorm:"unique;not null"`
    Password  string    `json:"-"`
    CreatedAt time.Time
    UpdatedAt time.Time

    // Relasi: satu user punya banyak URL
    URLs []Url `gorm:"foreignKey:UserID"`
}