package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/psocietyyy/go-gin-kafka/service/auth_service/internal/dto"
	userError "github.com/psocietyyy/go-gin-kafka/service/auth_service/internal/errors"
	"github.com/psocietyyy/go-gin-kafka/service/auth_service/internal/service"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var payload dto.LoginRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	token, err := h.authService.Login(payload.Email, payload.Password)
	if err != nil {
		if err == userError.ErrUserNotFound {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}
		if err == userError.ErrInvalidPassword {
			c.JSON(401, gin.H{"error": err.Error()})
			return
		}
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"token": token})
}

func (h *AuthHandler) Register(c *gin.Context) {
	var payload dto.CreateUserRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err := h.authService.Register(&payload)
	if err != nil {
		if err == userError.ErrEmailAlreadyExists {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, gin.H{"message": "User created successfully"})
}
