package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/psocietyyy/go-gin-kafka/service/auth_service/internal/config"
	"github.com/psocietyyy/go-gin-kafka/service/auth_service/internal/handler"
	authHandler "github.com/psocietyyy/go-gin-kafka/service/auth_service/internal/handler"
	"github.com/psocietyyy/go-gin-kafka/service/auth_service/internal/repository"
	"github.com/psocietyyy/go-gin-kafka/service/auth_service/internal/routes"
	"github.com/psocietyyy/go-gin-kafka/service/auth_service/internal/service"
	authService "github.com/psocietyyy/go-gin-kafka/service/auth_service/internal/service"
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

	repo := repository.NewUserRepository(config.DB)
	service := service.NewUserService(repo)
	handler := handler.NewUserHandler(service)

	authServices := authService.NewAuthService(repo)
	authHandlers := authHandler.NewAuthHandler(authServices)

	// Register routes
	routes.UserRoutes(r, handler)
	routes.AuthRoutes(r, authHandlers)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
