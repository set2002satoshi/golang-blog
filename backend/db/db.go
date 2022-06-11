package db

import (
	"log"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"github.com/set2002satoshi/golang-blog/model"
)


func DBInit() {
	DB := ConnectDB()
	DB.AutoMigrate(&model.CustomerInfo{})
	DB.AutoMigrate(&model.Customer{})
	DB.AutoMigrate(&model.Blog{})
	DB.AutoMigrate(&model.Tag{})
}

func ConnectDB() *gorm.DB {
	dsn := "docker:pass@tcp(app-db:3306)/main?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("OpenDB failed:", err)
	}
	return db
}

