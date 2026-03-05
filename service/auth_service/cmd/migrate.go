package main

import (
	"log"

	"github.com/psocietyyy/go-gin-kafka/service/auth_service/internal/config"
	"github.com/psocietyyy/go-gin-kafka/service/auth_service/internal/model"
)

func main() {
	// Connect To Database
	config.ConnectToDB()

	// Auto Migrate
	config.DB.AutoMigrate(&model.User{})

	log.Println("Migration User completed")
}