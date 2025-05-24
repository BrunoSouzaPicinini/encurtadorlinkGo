package main

import (
	"encurtadorLink/src/controller"
	"encurtadorLink/src/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	repository.Connect()

	router := gin.Default()

	router.POST("/shortener", controller.Create)
	router.GET("/r/:slug", controller.RedirectBySlug)

	router.Run(":8080") // Start the server on port 8080
}
