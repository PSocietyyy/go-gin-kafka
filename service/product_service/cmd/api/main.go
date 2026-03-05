package api

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/psocietyyy/go-gin-kafka/service/product_service/internal/config"
	"github.com/psocietyyy/go-gin-kafka/service/product_service/internal/handler"
	"github.com/psocietyyy/go-gin-kafka/service/product_service/internal/repository"
	"github.com/psocietyyy/go-gin-kafka/service/product_service/internal/routes"
	"github.com/psocietyyy/go-gin-kafka/service/product_service/internal/service"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Connect to database
	config.ConnectToDB()

	// Initialize repository
	trainRepo := repository.NewTrainRepository(config.DB)

	// Initialize service
	trainService := service.NewTrainService(trainRepo)

	// Initialize handler
	trainHandler := handler.NewTrainHandler(trainService)

	// Initialize router
	app := gin.Default()

	// Register routes
	routes.TrainRoutes(app, trainHandler)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := app.Run(":" + port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}