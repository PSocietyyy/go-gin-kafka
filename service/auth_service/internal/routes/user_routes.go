package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/psocietyyy/go-gin-kafka/service/auth_service/internal/handler"
)

func UserRoutes(router *gin.Engine, handler *handler.UserHandler) {
	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/", handler.CreateUser)
		userRoutes.GET("/", handler.FindAll)
		userRoutes.GET("/:id", handler.FindByID)
		userRoutes.PUT("/:id", handler.UpdateUser)
		userRoutes.DELETE("/:id", handler.DeleteUser)
	}
}