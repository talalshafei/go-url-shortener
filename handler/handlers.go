package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/talalshafei/go-url-shortener/shortener"
	"github.com/talalshafei/go-url-shortener/store"
)

type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
	UserId  string `json:"user_id" binding:"required"`
}

func CreateShortUrl(ctx *gin.Context) {
	var creationRequest UrlCreationRequest
	if err := ctx.ShouldBindJSON(&creationRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	shortUrl := shortener.GenerateShortLink(creationRequest.LongUrl, creationRequest.UserId)
	store.SaveUrlMapping(shortUrl, creationRequest.LongUrl, creationRequest.UserId)

	host := "http:/localhost:9808/"
	ctx.JSON(200, gin.H{
		"message":  "short url created successfully",
		"shortUrl": host + shortUrl,
	})
}

func HandleShortUrlRedirect(ctx *gin.Context) {
	shortUrl := ctx.Param("shortUrl")
	initialUrl := store.RetrieveInitialUrl(shortUrl)
	ctx.Redirect(302, initialUrl)
}
