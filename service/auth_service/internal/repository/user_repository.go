package repository

import (
	userError "github.com/psocietyyy/go-gin-kafka/service/auth_service/internal/errors"
	"github.com/psocietyyy/go-gin-kafka/service/auth_service/internal/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Get All user
func (r *UserRepository) FindAll() ([]model.User, error) {
	var users []model.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// Get User By ID
func (r *UserRepository) FindByID(id uint) (*model.User, error) {
	var user model.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, userError.ErrUserNotFound
	}
	return &user, nil
}

// Create User
func (r *UserRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

// Update User
func (r *UserRepository) Update(user *model.User) error {
	return r.db.Save(user).Error
}

// Delete User
func (r *UserRepository) Delete(id uint) error {
	return r.db.Delete(&model.User{}, id).Error
}

// Get User By Email
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, userError.ErrEmailAlreadyExists
	}
	return &user, nil
}
