package domain

import (
	"gorm.io/gorm"
	"time"
)

type Review struct {
	gorm.Model
	ID        uint
	Subject   string // 주제
	CreatedAt time.Time
	UpdatedAt time.Time
}
