package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"shorturl/database"
	"shorturl/models"
	"time"

	"github.com/gin-gonic/gin"
)

// 初始化伺服器和資料庫
func main() {
	router := gin.Default()

	// 初始化資料庫
	database.ConnectDatabase()

	// 路由設定
	router.POST("/shorten", shortenURL)
	router.GET("/:shortcode", redirectURL)

	// 啟動伺服器
	router.Run(":8080")
}

// 短網址生成邏輯
func shortenURL(c *gin.Context) {
	var request struct {
		LongURL string `json:"long_url" binding:"required"`
	}
	fmt.Println("request:", request)

	// 從請求中獲取長網址
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// 生成短碼
	shortcode := generateShortCode()

	// 將短碼與長網址儲存到資料庫
	urlMapping := models.URLMapping{
		ShortCode: shortcode,
		LongURL:   request.LongURL,
	}
	database.DB.Create(&urlMapping)
	port := "9000"
	// 回傳短網址
	c.JSON(http.StatusOK, gin.H{
		"shortcode": shortcode,
		"short_url": fmt.Sprintf("http://localhost:%s/%s", port, shortcode),
	})
}

// 短碼重定向邏輯
func redirectURL(c *gin.Context) {
	shortcode := c.Param("shortcode")

	// 從資料庫中查找對應的長網址
	var urlMapping models.URLMapping
	if err := database.DB.Where("short_code = ?", shortcode).First(&urlMapping).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	// 重定向到長網址
	c.Redirect(http.StatusFound, urlMapping.LongURL)
}

// 短碼生成器
func generateShortCode() string {
	rand.Seed(time.Now().UnixNano())
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	shortCode := make([]byte, 6)
	for i := range shortCode {
		shortCode[i] = charset[rand.Intn(len(charset))]
	}
	return string(shortCode)
}
