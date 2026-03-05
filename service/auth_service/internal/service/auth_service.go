package service

import (
	"github.com/psocietyyy/go-gin-kafka/service/auth_service/internal/dto"
	userError "github.com/psocietyyy/go-gin-kafka/service/auth_service/internal/errors"
	"github.com/psocietyyy/go-gin-kafka/service/auth_service/internal/model"
	"github.com/psocietyyy/go-gin-kafka/service/auth_service/internal/repository"
	"github.com/psocietyyy/go-gin-kafka/service/auth_service/internal/utils"
)

type AuthService struct {
	userRepo *repository.UserRepository
}

func NewAuthService(userRepo *repository.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (s *AuthService) Login(email string, password string) (string, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return "", userError.ErrUserNotFound
	}
	if !utils.ComparePassword(password, user.Password) {
		return "", userError.ErrInvalidPassword
	}
	return utils.GenerateJWT(user.ID, user.Name, user.Email)
}

func (s *AuthService) Register(payload *dto.CreateUserRequest) (error) {
	var user model.User
	user.Name = payload.Name
	user.Email = payload.Email
	user.Password = utils.HashPassword(payload.Password)
	return s.userRepo.Create(&user)
}