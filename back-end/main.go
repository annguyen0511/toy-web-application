package main

import (
	"backend/models"
	"backend/routes"
	"log"
	"os"

	"github.com/gin-contrib/cors"
)

func main() {
	models.Connect()

	r := routes.SetupRouter(models.GlobalDB)

	// CORS configuration
	corsConfig := cors.Config{
		AllowAllOrigins: true, // Allow all origins
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Content-Type", "Authorization"},
	}

	r.Use(cors.New(corsConfig)) // Use CORS middleware

	port := os.Getenv("PORT") // Get the port from the environment variable
	if port == "" {
		port = "8080" // Default to 8080 if not set
	}

	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}
