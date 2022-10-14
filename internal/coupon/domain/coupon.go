package domain

import (
	"gorm.io/gorm"
	"time"
)

type Coupon struct {
	gorm.Model
	ID        uint
	Title     string
	Duration  time.Duration
	Price     int
	CreatedAt time.Time
	UpdatedAt time.Time
}
