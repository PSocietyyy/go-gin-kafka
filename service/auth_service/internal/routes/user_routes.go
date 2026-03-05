package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/psocietyyy/go-gin-kafka/service/auth_service/internal/config"
	"github.com/psocietyyy/go-gin-kafka/service/auth_service/internal/handler"
	"github.com/psocietyyy/go-gin-kafka/service/auth_service/internal/repository"
	"github.com/psocietyyy/go-gin-kafka/service/auth_service/internal/service"
)

func UserRoutes(router *gin.Engine) {
	repo := repository.NewUserRepository(config.DB)
	service := service.NewUserService(repo)
	handler := handler.NewUserHandler(service)

	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/", handler.CreateUser)
		userRoutes.GET("/", handler.FindAll)
		userRoutes.GET("/:id", handler.FindByID)
		userRoutes.PUT("/:id", handler.UpdateUser)
		userRoutes.DELETE("/:id", handler.DeleteUser)
	}
}