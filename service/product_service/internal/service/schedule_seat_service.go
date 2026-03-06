package service

import (
	trainError "github.com/psocietyyy/go-gin-kafka/service/product_service/internal/errors"
	"github.com/psocietyyy/go-gin-kafka/service/product_service/internal/model"
	"github.com/psocietyyy/go-gin-kafka/service/product_service/internal/repository"
)

type ScheduleSeatService struct {
	scheduleSeatRepo *repository.ScheduleSeatRepository
}

func NewScheduleSeatService(scheduleSeatRepo *repository.ScheduleSeatRepository) *ScheduleSeatService {
	return &ScheduleSeatService{scheduleSeatRepo: scheduleSeatRepo}
}

func (s *ScheduleSeatService) FindAll() []model.ScheduleSeat {
	return s.scheduleSeatRepo.FindAll()
}

func (s *ScheduleSeatService) FindByID(id uint) (model.ScheduleSeat, error) {
	return s.scheduleSeatRepo.FindByID(id)
}

func (s *ScheduleSeatService) FindByScheduleID(scheduleID uint) []model.ScheduleSeat {
	return s.scheduleSeatRepo.FindByScheduleID(scheduleID)
}

func (s *ScheduleSeatService) Create(seat model.ScheduleSeat) (model.ScheduleSeat, error) {
	return s.scheduleSeatRepo.Create(seat)
}

// BookSeat changes the status of an available seat to booked
func (s *ScheduleSeatService) BookSeat(id uint) error {
	seat, err := s.scheduleSeatRepo.FindByID(id)
	if err != nil {
		return err
	}

	if seat.Status == model.SeatStatusBooked {
		return trainError.ErrSeatAlreadyBooked
	}

	return s.scheduleSeatRepo.UpdateStatus(id, model.SeatStatusBooked)
}

// CancelSeat changes the status of a booked seat back to available
func (s *ScheduleSeatService) CancelSeat(id uint) error {
	seat, err := s.scheduleSeatRepo.FindByID(id)
	if err != nil {
		return err
	}

	// Maybe add verification to make sure it's actually booked before canceling
	if seat.Status == model.SeatStatusAvailable {
		// Could just return nil or return an error depending on business logic
		return nil
	}

	return s.scheduleSeatRepo.UpdateStatus(id, model.SeatStatusAvailable)
}
