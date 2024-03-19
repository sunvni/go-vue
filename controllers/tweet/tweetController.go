package tweet

import (
	"govue/models/db"

	"github.com/gin-gonic/gin"
)

func GetTweets(c *gin.Context) {
	db.Init()
	tweets := db.GetAll()

	c.JSON(200, gin.H{
		"tweets": tweets,
	})
}

func AddTweet(c *gin.Context) {
	db.Init()
	// Get the JSON body from the request
	var content struct {
		Title   string `json:"title" binding:"required"`
		Content string `json:"content" binding:"required"`
	}
	c.Bind(&content)

	db.Insert(content.Content, content.Title)

	c.JSON(200, gin.H{
		"message": "success",
	})
}