package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/psocietyyy/go-gin-kafka/service/auth_service/internal/config"
	"github.com/psocietyyy/go-gin-kafka/service/auth_service/internal/routes"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Initialize database
	config.ConnectToDB()

	// Initialize router
	r := gin.Default()

	// Register routes
	routes.UserRoutes(r)
	routes.AuthRoutes(r)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
