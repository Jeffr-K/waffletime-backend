package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
	"waffletime/internal/user/domain"
)

var Db *gorm.DB

func InitializedDatabase() {
	var err error
	dns := "root:root@tcp(127.0.0.1:3306)/wanted_db?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic("failed to connect database.")
	}

	Db.AutoMigrate(&domain.User{})

	mysqlDB, err := Db.DB()
	mysqlDB.SetMaxIdleConns(10)
	mysqlDB.SetMaxOpenConns(100)
	mysqlDB.SetConnMaxLifetime(time.Hour)

	fmt.Println("Connected MySQL Database.")
}
