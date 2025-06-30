package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/longkhan786/url-shortener/db"
	"github.com/longkhan786/url-shortener/models"
)

func ShortenURL(c *gin.Context) {

	database := db.DB

	var request struct {
		OriginalUrl string `json:"originalUrl" binding:"required,url"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid request format. Please provide a valid URL.",
		})
		return
	}

	shortCode := uuid.New().String()
	url := models.URL{
		ShortCode:   shortCode,
		OriginalURL: request.OriginalUrl,
	}

	if err := database.Create(&url).Error; err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to shorten URL. Please try again later.",
		})
		return
	}

	shortenedURL := "http://" + c.Request.Host + "/" + shortCode

	c.JSON(200, gin.H{
		"shortenedUrl": shortenedURL,
		"shortCode":    shortCode,
		"originalUrl":  request.OriginalUrl,
		"message":      "URL shortened successfully!",
	})
}

func RedirectURL(c *gin.Context) {
	database := db.DB

	shortCode := c.Param("shortURL")

	var url models.URL

	if err := database.Where("short_code = ?", shortCode).First(&url).Error; err != nil {
		c.JSON(404, gin.H{
			"error": "Short URL not found.",
		})
		return
	}

	if err := database.Model(&url).UpdateColumn("click_count", url.ClickCount+1).Error; err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to update click count. Please try again later.",
		})
		return
	}

	c.Redirect(302, url.OriginalURL)
}
