package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/psocietyyy/go-gin-kafka/service/product_service/internal/handler"
)

func ScheduleSeatRoutes(app *gin.Engine, scheduleSeatHandler *handler.ScheduleSeatHandler) {
	app.GET("/schedule_seats", scheduleSeatHandler.FindAll)
	app.GET("/schedule_seats/:id", scheduleSeatHandler.FindByID)
	app.GET("/schedules/:id/seats", scheduleSeatHandler.FindByScheduleID)
	app.POST("/schedule_seats", scheduleSeatHandler.Create)
	app.POST("/schedule_seats/:id/book", scheduleSeatHandler.BookSeat)
	app.POST("/schedule_seats/:id/cancel", scheduleSeatHandler.CancelSeat)
}
