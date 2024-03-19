package routes

import (
	homeController "govue/controllers/home"
	"govue/controllers/tweet"

	"github.com/gin-gonic/gin"
)


func HandleApi(router *gin.Engine) {
	// Routes
	apiRoutes := router.Group("/api/v1")

	apiRoutes.Use(func (c *gin.Context) {
        c.Header("Content-Type", "application/json")
        c.Header("X-Custom-Header", "custom")
        c.Next()
    })

	apiRoutes.POST("/counter/increment", homeController.IncrementCounter)
	apiRoutes.POST("/counter/decrement", homeController.DecrementCounter)
	apiRoutes.GET("/tweets", tweet.GetTweets)
	apiRoutes.POST("/tweet", tweet.AddTweet)
}