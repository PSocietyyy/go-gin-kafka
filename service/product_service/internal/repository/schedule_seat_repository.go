package repository

import (
	"github.com/psocietyyy/go-gin-kafka/service/product_service/internal/errors"
	"github.com/psocietyyy/go-gin-kafka/service/product_service/internal/model"
	"gorm.io/gorm"
)

type ScheduleSeatRepository struct {
	db *gorm.DB
}

func NewScheduleSeatRepository(db *gorm.DB) *ScheduleSeatRepository {
	return &ScheduleSeatRepository{db: db}
}

func (r *ScheduleSeatRepository) FindAll() []model.ScheduleSeat {
	var seats []model.ScheduleSeat
	if err := r.db.Preload("TrainSeat").Preload("Schedule").Preload("Schedule.Train").Find(&seats).Error; err != nil {
		return nil
	}
	return seats
}

func (r *ScheduleSeatRepository) FindByID(id uint) (model.ScheduleSeat, error) {
	var seat model.ScheduleSeat
	if err := r.db.Preload("TrainSeat").Preload("Schedule").Preload("Schedule.Train").First(&seat, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return model.ScheduleSeat{}, errors.ErrScheduleSeatNotFound
		}
		return model.ScheduleSeat{}, err
	}
	return seat, nil
}

func (r *ScheduleSeatRepository) FindByScheduleID(scheduleID uint) []model.ScheduleSeat {
	var seats []model.ScheduleSeat
	if err := r.db.Where("schedule_id = ?", scheduleID).Preload("TrainSeat").Preload("Schedule").Preload("Schedule.Train").Find(&seats).Error; err != nil {
		return nil
	}
	return seats
}

func (r *ScheduleSeatRepository) Create(seat model.ScheduleSeat) (model.ScheduleSeat, error) {
	if err := r.db.Create(&seat).Error; err != nil {
		return model.ScheduleSeat{}, err
	}
	return seat, nil
}

func (r *ScheduleSeatRepository) UpdateStatus(id uint, status model.SeatStatus) error {
	result := r.db.Model(&model.ScheduleSeat{}).Where("id = ?", id).Update("status", status)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.ErrScheduleSeatNotFound
	}
	return nil
}
