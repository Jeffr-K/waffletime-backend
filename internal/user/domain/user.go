package domain

import "time"

type User struct {
	ID       uint      `gorm:"primaryKey"`
	Username string    `gorm:"not null"`
	Password string    `gorm:"not null"`
	Email    string    `gorm:"unique"`
	Address  string    `gorm:"not null"`
	CreateAt time.Time `gorm:"autoCreateTime:milli"`
	UpdateAt time.Time `gorm:"autoUpdateTime:milli"`
	DeleteAt time.Time `gorm:"index"`
}

func (u User) Register(user User) User {
	return User{}
}

func (u User) Dropdown(user User) User {
	return User{}
}

func (u User) updateUser(user User) User {
	return User{}
}

func (u User) getUser(user User) User {
	return User{}
}
