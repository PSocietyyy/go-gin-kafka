package main

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
	trainSeatRepo := repository.NewTrainSeatRepository(config.DB)
	scheduleRepo := repository.NewScheduleRepository(config.DB)
	scheduleSeatRepo := repository.NewScheduleSeatRepository(config.DB)

	// Initialize service
	trainService := service.NewTrainService(trainRepo)
	trainSeatService := service.NewTrainSeatService(trainSeatRepo, trainRepo)
	scheduleService := service.NewScheduleService(scheduleRepo)
	scheduleSeatService := service.NewScheduleSeatService(scheduleSeatRepo)

	// Initialize handler
	trainHandler := handler.NewTrainHandler(trainService)
	trainSeatHandler := handler.NewTrainSeatHandler(trainSeatService)
	scheduleHandler := handler.NewScheduleHandler(scheduleService)
	scheduleSeatHandler := handler.NewScheduleSeatHandler(scheduleSeatService)

	// Initialize router
	app := gin.Default()

	// Register routes
	routes.TrainRoutes(app, trainHandler)
	routes.TrainSeatRoutes(app, trainSeatHandler)
	routes.ScheduleRoutes(app, scheduleHandler)
	routes.ScheduleSeatRoutes(app, scheduleSeatHandler)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	if err := app.Run(":" + port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}