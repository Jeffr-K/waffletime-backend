package domain

import (
	"gorm.io/gorm"
	"time"
)

type Review struct {
	gorm.Model
	ID        uint
	Title     string // 주제
	Content   string // 내용
	Photo     string
	Star      int
	CreatedAt time.Time
	UpdatedAt time.Time
	UserId    string
}
