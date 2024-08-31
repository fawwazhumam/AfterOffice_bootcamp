package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {

	const MYSQL = "root:@tcp(127.0.0.1:3306)/crud_go?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := MYSQL
	var err error

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Cannot connect to database: %v", err))
	}
	fmt.Println("Connected to database")
}