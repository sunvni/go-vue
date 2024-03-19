package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleWeb(router *gin.Engine) {
	router.LoadHTMLGlob("public/*.html")
	router.Static("/assets", "./public/assets")
	router.Static("/favicon.ico", "./public/favicon.ico")

	router.GET("/hello", func(c *gin.Context) {
		c.HTML(http.StatusOK, "hello.html", nil)
	})
}