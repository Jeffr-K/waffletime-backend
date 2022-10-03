package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jeffr-k/waffletime/internal/auth"
	"github.com/jeffr-k/waffletime/internal/user/presentor"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type Bootstrap struct{}

func (b Bootstrap) initializeRouter(router *gin.Engine) {
	user := presentor.Router{}
	auth := auth.Router{}

	user.Routes(router)
	auth.Routes(router)
}

func (b Bootstrap) initializedDatabase() {
	dns := "root:root@tcp(127.0.0.1:3306)/wanted_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic("failed to connect database.")
	}
	mysqlDB, err := db.DB()
	mysqlDB.SetMaxIdleConns(10)
	mysqlDB.SetMaxOpenConns(100)
	mysqlDB.SetConnMaxLifetime(time.Hour)

	fmt.Println("Connected MySQL Database.")
}
