package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DB database object
var DB *gorm.DB

// StartDbConnection : connect to database
func StartDbConnection() {
	// database init
	database, err := gorm.Open(sqlite.Open("test1.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	database.AutoMigrate(&User{})

	DB = database
}
