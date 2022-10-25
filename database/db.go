package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// github.com/denisenkom/go-mssqldb
//   dsn := "sqlserver://gorm:LoremIpsum86@localhost:9930?database=gorm"
//   db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

// github.com/mattn/go-sqlite3
//db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

var Db *gorm.DB

func InitDb() *gorm.DB { // OOP constructor
	Db = connectDB()
	return Db
}

func connectDB() *gorm.DB {
	dsn := "host=localhost user=TestDB password=654321 dbname=shoppingcart port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error...")
		return nil
	}
	return db
}
