package common

import (
	"dyk/model"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	username := "root"
	password := "Admwork0620"
	host := "localhost"
	port := "3306"
	database := "gin"
	charset := "utf8mb4"
	// driverName := "mysql"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		database,
		charset,
	)

	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	}
	// 自动创建数据表
	DB.AutoMigrate(&model.User{})
	return DB
}

func GetDB() *gorm.DB {
	return DB
}
