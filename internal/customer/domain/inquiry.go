package domain

import (
	"gorm.io/gorm"
	"time"
)

type Inquiry struct {
	gorm.Model
	ID         uint
	Content    string // 문의 내용
	CreatedAt  time.Time
	UpdatedAt  time.Time
	ConsumerId string // 소비자
	ProducerId string // 생산자
}
