package service

import (
	"github.com/psocietyyy/go-gin-kafka/service/auth_service/internal/dto"
	userError "github.com/psocietyyy/go-gin-kafka/service/auth_service/internal/errors"
	"github.com/psocietyyy/go-gin-kafka/service/auth_service/internal/model"
	"github.com/psocietyyy/go-gin-kafka/service/auth_service/internal/repository"
	"github.com/psocietyyy/go-gin-kafka/service/auth_service/internal/utils"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// Create User
func (s *UserService) CreateUser(payload *dto.CreateUserRequest) error {
	var user model.User
	user.Name = payload.Name
	user.Email = payload.Email
	user.Password = utils.HashPassword(payload.Password)
	return s.repo.Create(&user)
}

// Get All User
func (s *UserService) FindAll() ([]model.User, error) {
	return s.repo.FindAll()
}

// Get User By ID
func (s *UserService) FindByID(id uint) (*model.User, error) {
	return s.repo.FindByID(id)
}

// Update User
func (s *UserService) UpdateUser(payload *dto.UpdateUserRequest, id uint) error {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return userError.ErrUserNotFound
	}
	if payload.Name != "" {
		user.Name = payload.Name
	}
	if payload.Email != "" {
		// Check if email already exists
		existingUser, err := s.repo.FindByEmail(payload.Email)
		if err != nil {
			return err
		}
		if existingUser != nil && existingUser.ID != id {
			return userError.ErrEmailAlreadyExists
		}
		user.Email = payload.Email
	}
	if payload.Password != "" {
		user.Password = utils.HashPassword(payload.Password)
	}
	return s.repo.Update(user)
}

// Delete User
func (s *UserService) DeleteUser(id uint) error {
	return s.repo.Delete(id)
}

// Get User By Email
func (s *UserService) FindByEmail(email string) (*model.User, error) {
	return s.repo.FindByEmail(email)
}