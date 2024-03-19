package routes

import (
	homeController "govue/controllers/home"
	"govue/controllers/tweet"

	"github.com/gin-gonic/gin"
)


func Handle(router *gin.Engine) {

	prefix := "/api/v1"

	

	router.LoadHTMLGlob("public/*.html")
	router.Static("/assets", "./public/assets")
	router.Static("/favicon.ico", "./public/favicon.ico")

	// Routes
	router.POST(prefix + "/counter/increment", homeController.IncrementCounter)
	router.POST(prefix + "/counter/decrement", homeController.DecrementCounter)
	router.GET(prefix + "/tweets", tweet.GetTweets)
	router.POST(prefix + "/tweet", tweet.AddTweet)
}