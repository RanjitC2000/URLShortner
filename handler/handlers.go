package handler

import (
	"net/http"

	"github.com/RanjitC2000/go-url-shortener/shortner"
	"github.com/RanjitC2000/go-url-shortener/store"
	"github.com/gin-gonic/gin"
)

type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
	UserID  string `json:"user_id" binding:"required"`
}

func CreateShortUrl(c *gin.Context) {
	var creationRequest UrlCreationRequest
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortUrl := shortner.GenerateShortLink(creationRequest.LongUrl, creationRequest.UserID)
	store.SaveUrlMapping(shortUrl, creationRequest.LongUrl, creationRequest.UserID)

	host := "http://localhost:9808/"
	c.JSON(http.StatusOK, gin.H{
		"message":   "Short URL created successfully",
		"short_url": host + shortUrl,
	})

}

func HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	initialUrl := store.RetrieveInitialUrl(shortUrl)
	c.Redirect(302, initialUrl)
}
