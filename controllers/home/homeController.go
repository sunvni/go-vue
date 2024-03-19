package homeController

import (
	"math/rand"

	"github.com/gin-gonic/gin"
)

func IncrementCounter(c *gin.Context) {
	c.JSON(200, gin.H{
		"count": rand.Intn(100),
	})
}

func DecrementCounter(c *gin.Context) {
	c.JSON(200, gin.H{
		"count": rand.Intn(100),
	})
}
