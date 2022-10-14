package domain

import (
	"gorm.io/gorm"
	"time"
)

type Cart struct {
	gorm.Model
	ID     uint
	Amount int // 장바구니 상품 수량
	//Products        []Product.Product // 장바구니 상품 목록
	TotalPrice      int       // 장바구니 총 상품 가격
	DeliveryFee     int       // 장바구니 배달 금액
	TotalOrderPrice int       // 장바구니 총 주문가격
	CreatedAt       time.Time // 장바구니 생성일
	UpdatedAt       time.Time // 장바구니 변경일
	UserId          string    // 장바구니 소비자
}
