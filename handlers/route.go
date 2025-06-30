package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Route(router *gin.Engine) {

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "URL Shortener API is running!",
		})
	})

	router.POST("/shorten", ShortenURL)
	router.GET("/:shortURL", RedirectURL)
}
