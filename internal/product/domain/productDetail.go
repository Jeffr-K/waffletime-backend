package domain

import (
	"gorm.io/gorm"
	"time"
)

type ProductDetail struct {
	gorm.Model
	ID        string
	Number    string    // 상품 번호
	Category  string    // 상품 카테고리
	Name      string    // 상품 이름
	Price     string    // 상품 가격
	Weight    string    // 상품 중량
	Photo     string    // 상품 사진
	Content   string    // 상품 내용
	Thumbnail string    // 상품 대표사진
	CreatedAt time.Time // 등록된 일자
	UpdatedAt time.Time // 변경된 일자
	UserId    string    // 판매자
}
