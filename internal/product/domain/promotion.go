package domain

import "time"

type Promotion struct {
	ID       uint
	Name     string    // 할인명
	Duration time.Time // 할인기간
	Price    int       // 할인금액
	Rate     float64   // 할인비율
	Range    string    // 할인적용범위
	User     string    // 할인대상
	Active   bool      // 집행여부
}
