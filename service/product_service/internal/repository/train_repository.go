package repository

import (
	"github.com/psocietyyy/go-gin-kafka/service/product_service/internal/errors"
	"github.com/psocietyyy/go-gin-kafka/service/product_service/internal/model"
	"gorm.io/gorm"
)

type TrainRepository struct {
	db *gorm.DB
}

func NewTrainRepository(db *gorm.DB) *TrainRepository {
	return &TrainRepository{db: db}
}

// Get All Trains
func (r *TrainRepository) FindAll() ([]model.Train) {
	var trains []model.Train
	if err := r.db.Find(&trains).Error; err != nil {
		return nil
	}
	return trains
}

// Get Train By ID
func (r *TrainRepository) FindByID(id uint) (model.Train, error) {
	var train model.Train
	if err := r.db.First(&train, id).Error; err != nil {
		// Check if the error is "record not found"
		if err == gorm.ErrRecordNotFound {
			return model.Train{}, errors.ErrTrainNotFound
		}
		return model.Train{}, err
	}
	return train, nil
}

// Create Train
func (r *TrainRepository) Create(train model.Train) (model.Train, error) {
	if err := r.db.Create(&train).Error; err != nil {
		return model.Train{}, err
	}
	return train, nil
}

// Update Train
func (r *TrainRepository) Update(train model.Train) (model.Train, error) {
	if err := r.db.Save(&train).Error; err != nil {
		return model.Train{}, err
	}
	return train, nil
}

// Delete Train
func (r *TrainRepository) Delete(id uint) error {
	if err := r.db.Delete(&model.Train{}, id).Error; err != nil {
		return err
	}
	return nil
}

// Find By Code
func (r *TrainRepository) FindByCode(code string) (model.Train, error) {
	var train model.Train
	if err := r.db.Where("code = ?", code).First(&train).Error; err != nil {
		// Check if the error is "record not found"
		if err == gorm.ErrRecordNotFound {
			return model.Train{}, errors.ErrTrainCodeAlreadyExists
		}
		return model.Train{}, err
	}
	return train, nil
}
