package domain

import (
	"gorm.io/gorm"
	"time"
)

type Policy struct {
	gorm.Model
	ID        uint
	Type      string  // 포인트 적립 타입
	Rate      float32 // 구매시 할인 적용 비율
	CreatedAt time.Time
	UpdatedAt time.Time
}
