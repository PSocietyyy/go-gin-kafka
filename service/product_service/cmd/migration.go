package main

import (
	"github.com/psocietyyy/go-gin-kafka/service/product_service/internal/config"
	"github.com/psocietyyy/go-gin-kafka/service/product_service/internal/model"
)

func main() {
	config.ConnectToDB()
	config.DB.AutoMigrate(
		&model.Train{},
		&model.TrainSeat{},
		&model.Schedule{},
		&model.ScheduleSeat{},
	)
}