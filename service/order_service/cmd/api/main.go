package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/psocietyyy/go-gin-kafka/service/order_service/internal/config"
	"github.com/psocietyyy/go-gin-kafka/service/order_service/internal/handler"
	"github.com/psocietyyy/go-gin-kafka/service/order_service/internal/repository"
	"github.com/psocietyyy/go-gin-kafka/service/order_service/internal/routes"
	"github.com/psocietyyy/go-gin-kafka/service/order_service/internal/service"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Connect to database
	config.ConnectToDB()

	// Initialize repository
	orderRepo := repository.NewOrderRepository(config.DB)

	// Initialize service
	orderService := service.NewOrderService(orderRepo)

	// Initialize handler
	orderHandler := handler.NewOrderHandler(orderService)

	// Initialize router
	app := gin.Default()

	// Register routes
	routes.OrderRoutes(app, orderHandler)

	// Start server on port 8082
	port := os.Getenv("PORT")
	if port == "" {
		port = "8082" // Use 8082 for Order Service
	}

	if err := app.Run(":" + port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
