package domain

import (
	"gorm.io/gorm"
	"time"
)

type Point struct {
	gorm.Model
	ID        uint
	Title     string
	Duration  time.Time
	Price     int
	CreatedAt time.Time
	UpdatedAt time.Time
	SellerId  string // 포인트 발행자
}
