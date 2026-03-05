package service

import (
	"github.com/psocietyyy/go-gin-kafka/service/product_service/internal/model"
	"github.com/psocietyyy/go-gin-kafka/service/product_service/internal/repository"
)

type ScheduleService struct {
	scheduleRepo *repository.ScheduleRepository
}

func NewScheduleService(scheduleRepo *repository.ScheduleRepository) *ScheduleService {
	return &ScheduleService{scheduleRepo: scheduleRepo}
}

func (s *ScheduleService) FindAll() []model.Schedule {
	return s.scheduleRepo.FindAll()
}

func (s *ScheduleService) FindByID(id uint) (model.Schedule, error) {
	return s.scheduleRepo.FindByID(id)
}

func (s *ScheduleService) Create(schedule model.Schedule) (model.Schedule, error) {
	return s.scheduleRepo.Create(schedule)
}

func (s *ScheduleService) Update(schedule model.Schedule) (model.Schedule, error) {
	_, err := s.scheduleRepo.FindByID(schedule.ID)
	if err != nil {
		return model.Schedule{}, err
	}
	return s.scheduleRepo.Update(schedule)
}

func (s *ScheduleService) Delete(id uint) error {
	_, err := s.scheduleRepo.FindByID(id)
	if err != nil {
		return err
	}
	return s.scheduleRepo.Delete(id)
}
