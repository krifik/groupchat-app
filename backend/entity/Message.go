package entity

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	UserID int
	// User    User   `gorm:"foreignKey:UserID"`
	Content string `gorm:"size:10000"`
}
