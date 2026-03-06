package repository

import (
	"github.com/psocietyyy/go-gin-kafka/service/product_service/internal/errors"
	"github.com/psocietyyy/go-gin-kafka/service/product_service/internal/model"
	"gorm.io/gorm"
)

type TrainSeatRepository struct {
	db *gorm.DB
}

func NewTrainSeatRepository(db *gorm.DB) *TrainSeatRepository {
	return &TrainSeatRepository{db: db}
}

func (r *TrainSeatRepository) FindAll() []model.TrainSeat {
	var seats []model.TrainSeat
	if err := r.db.Find(&seats).Error; err != nil {
		return nil
	}
	return seats
}

func (r *TrainSeatRepository) FindByID(id uint) (model.TrainSeat, error) {
	var seat model.TrainSeat
	if err := r.db.First(&seat, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return model.TrainSeat{}, errors.ErrTrainSeatNotFound
		}
		return model.TrainSeat{}, err
	}
	return seat, nil
}

func (r *TrainSeatRepository) FindByTrainID(trainID uint) []model.TrainSeat {
	var seats []model.TrainSeat
	if err := r.db.Where("train_id = ?", trainID).Find(&seats).Error; err != nil {
		return nil
	}
	return seats
}

func (r *TrainSeatRepository) Create(seat model.TrainSeat) (model.TrainSeat, error) {
	if err := r.db.Create(&seat).Error; err != nil {
		return model.TrainSeat{}, err
	}
	return seat, nil
}

func (r *TrainSeatRepository) Update(seat model.TrainSeat) (model.TrainSeat, error) {
	if err := r.db.Save(&seat).Error; err != nil {
		return model.TrainSeat{}, err
	}
	return seat, nil
}

func (r *TrainSeatRepository) Delete(id uint) error {
	if err := r.db.Delete(&model.TrainSeat{}, id).Error; err != nil {
		return err
	}
	return nil
}
