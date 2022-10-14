package domain

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"type:varchar(50)"`
	Password  string `gorm:"type:varchar(150)"`
	Email     string `gorm:"type:varchar(100);unique"`
	Address   string `gorm:"type:varchar(100)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
