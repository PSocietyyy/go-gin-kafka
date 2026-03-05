package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/psocietyyy/go-gin-kafka/service/order_service/internal/handler"
)

func OrderRoutes(app *gin.Engine, orderHandler *handler.OrderHandler) {
	app.GET("/orders", orderHandler.FindAll)
	app.POST("/orders", orderHandler.CreateOrder)
	app.GET("/orders/:id", orderHandler.FindByID)
	app.PUT("/orders/:id", orderHandler.Update)
}
