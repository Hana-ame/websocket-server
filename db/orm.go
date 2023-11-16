package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db = initDB()
)

// init a db
func initDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})

	return db
}

func Init() {
	db = initDB()
}

func init() {
	Init()
}
