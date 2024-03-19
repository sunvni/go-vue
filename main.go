package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"

	// Import the controllers package
	"govue/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	gin.SetMode(os.Getenv("GIN_MODE"))

	r := gin.Default()

	routes.HandleApi(r)

	routes.HandleWeb(r)

	// Run the server on port 8080
	r.Run(":8080")
}
