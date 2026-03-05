package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/psocietyyy/go-gin-kafka/service/auth_service/internal/config"
	"github.com/psocietyyy/go-gin-kafka/service/auth_service/internal/handler"
	"github.com/psocietyyy/go-gin-kafka/service/auth_service/internal/repository"
	"github.com/psocietyyy/go-gin-kafka/service/auth_service/internal/service"
)

func AuthRoutes(router *gin.Engine) {
	repo := repository.NewUserRepository(config.DB)
	service := service.NewAuthService(repo)
	handler := handler.NewAuthHandler(service)
	authRoutes := router.Group("/auth")
	authRoutes.POST("/login", handler.Login)
	authRoutes.POST("/register", handler.Register)
}