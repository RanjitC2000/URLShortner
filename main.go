package main

import (
	"fmt"
	"path/filepath"

	"github.com/RanjitC2000/go-url-shortener/handler"
	"github.com/RanjitC2000/go-url-shortener/store"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Serve static files (HTML, JS, CSS)
	r.Static("/static", "./static")

	// Serve index.html as the homepage
	r.GET("/", func(c *gin.Context) {
		c.File(filepath.Join("static", "index.html"))
	})

	// API endpoints
	r.POST("/create-short-url", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})

	r.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c)
	})

	store.InitializeStore()

	err := r.Run(":9808")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}
