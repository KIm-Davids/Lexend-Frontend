package models

import "time"

type User struct {
	Email      string `gorm:"uniqueIndex;size:255" json:"email"` // Specify size to limit the index length
	Password   string `json:"password"`
	IsLoggedIn string
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
