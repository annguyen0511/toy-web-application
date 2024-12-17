package main

import (
	"backend/models"
	"backend/routes"
	"log"
	"os"
)

func main() {
	models.Connect()

	r := routes.SetupRouter(models.GlobalDB)

	port := os.Getenv("PORT") // Get the port from the environment variable
	if port == "" {
		port = "8080" // Default to 8080 if not set
	}

	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}
