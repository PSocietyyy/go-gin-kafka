package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/psocietyyy/go-gin-kafka/service/product_service/internal/handler"
)

func TrainRoutes(app *gin.Engine, trainHandler *handler.TrainHandler) {
	app.GET("/trains", trainHandler.FindAll)
	app.GET("/trains/:id", trainHandler.FindByID)
	app.POST("/trains", trainHandler.Create)
	app.PUT("/trains/:id", trainHandler.Update)
	app.DELETE("/trains/:id", trainHandler.Delete)
}