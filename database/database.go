package database

import (
	"log"
	"shorturl/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// 連接資料庫並創建 URL 對應表
func ConnectDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("shorturl.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	log.Println("Database connected")

	// 自動遷移（創建表）
	DB.AutoMigrate(&models.URLMapping{})
}
