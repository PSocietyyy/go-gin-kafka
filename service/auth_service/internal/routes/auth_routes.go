package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/psocietyyy/go-gin-kafka/service/auth_service/internal/handler"
)

func AuthRoutes(router *gin.Engine, handler *handler.AuthHandler) {
	authRoutes := router.Group("/auth")
	authRoutes.POST("/login", handler.Login)
	authRoutes.POST("/register", handler.Register)
}