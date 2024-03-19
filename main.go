package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"

	// Import the controllers package
	"govue/routes"
)

// (Assuming your API controllers are defined in a separate package)

// Replace with your actual API prefix (e.g., "/api/v1")
//const apiPrefix string = "/api/v1"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	gin.SetMode(os.Getenv("GIN_MODE"))

	r := gin.Default()

	routes.Handle(r)

	// Run the server on port 8080
	r.Run(":8080")

}
