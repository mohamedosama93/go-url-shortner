package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mohamedosama93/url-shortener/shortener"
	"github.com/mohamedosama93/url-shortener/store"
)

type UrlCreationRequest struct {
	Url string `json:"url" binding:"required"`
	UserId string `json:"user_id" binding:"required"`
}

func CreateShortUrl(c *gin.Context) {
	var creationRequest UrlCreationRequest;
	err := c.ShouldBindJSON(&creationRequest);
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return;
	}

	shortUrl := shortener.GenerateShortLink(creationRequest.Url, creationRequest.UserId)
	store.SaveUrl(shortUrl, creationRequest.Url, creationRequest.UserId)

	c.JSON(http.StatusOK, gin.H{
		"url": fmt.Sprintf("http://localhost:3000/%s", shortUrl),
		"message": "Success",
	})
}

func GetShortUrl(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	url := store.GetUrl(shortUrl)
	c.Redirect(http.StatusFound, url)
}