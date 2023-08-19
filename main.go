package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mohamedosama93/url-shortener/handlers"
	"github.com/mohamedosama93/url-shortener/store"
)

func main() {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "URL Shortener",
		})
	})

	r.POST("/create", func(ctx *gin.Context) {
		handlers.CreateShortUrl(ctx);
	})

	r.GET("/:shortUrl", func(ctx *gin.Context) {
		handlers.GetShortUrl(ctx);
	})

	store.InitializeStore();

	err := r.Run(":3000")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error %v", err))
	}
}