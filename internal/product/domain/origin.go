package domain

import (
	"gorm.io/gorm"
	"time"
)

type Origin struct {
	gorm.Model
	ID              uint
	Name            string    // 원산지명
	Duration        time.Time // 유통기한
	ProductNumber   string    // 상품번호
	ProductCategory string    // 상품카테고리
}
