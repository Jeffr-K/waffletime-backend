package domain

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	ID uint
}
