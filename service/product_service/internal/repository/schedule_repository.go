package repository

import (
	"github.com/psocietyyy/go-gin-kafka/service/product_service/internal/errors"
	"github.com/psocietyyy/go-gin-kafka/service/product_service/internal/model"
	"gorm.io/gorm"
)

type ScheduleRepository struct {
	db *gorm.DB
}

func NewScheduleRepository(db *gorm.DB) *ScheduleRepository {
	return &ScheduleRepository{db: db}
}

func (r *ScheduleRepository) FindAll() []model.Schedule {
	var schedules []model.Schedule
	if err := r.db.Preload("Train").Find(&schedules).Error; err != nil {
		return nil
	}
	return schedules
}

func (r *ScheduleRepository) FindByID(id uint) (model.Schedule, error) {
	var schedule model.Schedule
	if err := r.db.Preload("Train").First(&schedule, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return model.Schedule{}, errors.ErrScheduleNotFound
		}
		return model.Schedule{}, err
	}
	return schedule, nil
}

func (r *ScheduleRepository) Create(schedule model.Schedule) (model.Schedule, error) {
	if err := r.db.Create(&schedule).Error; err != nil {
		return model.Schedule{}, err
	}
	return schedule, nil
}

func (r *ScheduleRepository) Update(schedule model.Schedule) (model.Schedule, error) {
	if err := r.db.Save(&schedule).Error; err != nil {
		return model.Schedule{}, err
	}
	return schedule, nil
}

func (r *ScheduleRepository) Delete(id uint) error {
	if err := r.db.Delete(&model.Schedule{}, id).Error; err != nil {
		return err
	}
	return nil
}
