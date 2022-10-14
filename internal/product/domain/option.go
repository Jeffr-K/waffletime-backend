package domain

import "gorm.io/gorm"

type Option struct {
	gorm.Model
	ID    uint
	Name  string
	Value string
	Code  string
}
