package main

import (
	"log"

	"github.com/psocietyyy/go-gin-kafka/service/order_service/internal/config"
	"github.com/psocietyyy/go-gin-kafka/service/order_service/internal/model"
)

func main() {
	// Connect to database
	config.ConnectToDB()

	config.DB.AutoMigrate(&model.Order{})

	log.Println("Database migrated")
}