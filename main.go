package main

import (
	"github.com/gin-gonic/gin"
	"github.com/longkhan786/url-shortener/db"
	"github.com/longkhan786/url-shortener/handlers"
)

func main() {

	router := gin.Default()
	db.InitDB()

	handlers.Route(router)

	// Start server on port 8080
	router.Run(":8080")

}
