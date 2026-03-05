package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/psocietyyy/go-gin-kafka/service/product_service/internal/handler"
)

func TrainSeatRoutes(app *gin.Engine, trainSeatHandler *handler.TrainSeatHandler) {
	app.GET("/train_seats/:id", trainSeatHandler.FindByID)
	app.GET("/trains/:id/seats", trainSeatHandler.FindByTrainID)
	app.POST("/train_seats", trainSeatHandler.Create)
	app.PUT("/train_seats/:id", trainSeatHandler.Update)
	app.DELETE("/train_seats/:id", trainSeatHandler.Delete)
}
