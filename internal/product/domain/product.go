package domain

import (
	"gorm.io/gorm"
	"time"
	ReviewModule "waffletime/internal/review/domain"
	UserModule "waffletime/internal/user/domain"
)

type Product struct {
	gorm.Model
	Title         string    // 상품판매타이틀
	Option        []string  // 상품 옵션
	Star          float32   // 상품 별점
	CreatedAt     time.Time // 상품 게시 날짜
	UpdatedAt     time.Time // 상품 업데이트 날짜
	ProductDetail ProductDetail
	UserId        UserModule.User
	ReviewId      ReviewModule.Review
}

type ProductDetail struct {
	gorm.Model
	ID        string
	Number    string
	Category  string
	Name      string
	Price     string
	Photo     string
	CreatedAt time.Time
	UpdatedAt time.Time
	UserId    string
}
