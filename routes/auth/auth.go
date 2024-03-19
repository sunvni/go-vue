package auth

import (
	"github.com/gin-gonic/gin"
)

// Login is a function to handle the login route
func Login(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "login",
	})
}

// Register is a function to handle the register route
func Register(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "register",
	})
}