package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/psocietyyy/go-gin-kafka/service/product_service/internal/handler"
)

func ScheduleRoutes(app *gin.Engine, scheduleHandler *handler.ScheduleHandler) {
	app.GET("/schedules", scheduleHandler.FindAll)
	app.GET("/schedules/:id", scheduleHandler.FindByID)
	app.POST("/schedules", scheduleHandler.Create)
	app.PUT("/schedules/:id", scheduleHandler.Update)
	app.DELETE("/schedules/:id", scheduleHandler.Delete)
}
