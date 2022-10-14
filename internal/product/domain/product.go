package domain

import (
	"gorm.io/gorm"
	"time"
)

type Product struct {
	gorm.Model
	Title     string    // 상품판매타이틀
	Option    string    // 상품 옵션
	Star      float32   // 상품 별점
	Amount    int       // 상품 수량
	IsSoldOut bool      // 품절 여부
	CreatedAt time.Time // 상품 게시 날짜
	UpdatedAt time.Time // 상품 업데이트 날짜
	//ProductDetail ProductDetail
	//Origin        Origin
	//UserId        UserModule.User
	//ReviewId      ReviewModule.Review
}
