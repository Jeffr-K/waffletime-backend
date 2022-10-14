package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
	Cart "waffletime/internal/cart/domain"
	Coupon "waffletime/internal/coupon/domain"
	CustomerService "waffletime/internal/customer/domain"
	Point "waffletime/internal/point/domain"
	Product "waffletime/internal/product/domain"
	Review "waffletime/internal/review/domain"
	User "waffletime/internal/user/domain"
)

var Db *gorm.DB

func InitializedDatabase() {
	var err error
	dns := "root:root@tcp(127.0.0.1:3306)/wanted_db?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic("failed to connect database.")
	}

	Db.AutoMigrate(&User.User{})
	Db.AutoMigrate(&Product.Product{})
	Db.AutoMigrate(&Product.ProductDetail{})
	Db.AutoMigrate(&Product.Category{})
	Db.AutoMigrate(&Product.Option{})
	Db.AutoMigrate(&Product.Origin{})
	Db.AutoMigrate(&Product.Promotion{})
	Db.AutoMigrate(&Cart.Cart{})
	Db.AutoMigrate(&Review.Review{})
	Db.AutoMigrate(&Coupon.Coupon{})
	Db.AutoMigrate(&Point.Point{})
	Db.AutoMigrate(&Point.Policy{})
	Db.AutoMigrate(&CustomerService.Inquiry{})

	mysqlDB, err := Db.DB()
	mysqlDB.SetMaxIdleConns(10)
	mysqlDB.SetMaxOpenConns(100)
	mysqlDB.SetConnMaxLifetime(time.Hour)

	fmt.Println("Connected MySQL Database.")
}
