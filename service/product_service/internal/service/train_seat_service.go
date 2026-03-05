package service

import (
	trainError "github.com/psocietyyy/go-gin-kafka/service/product_service/internal/errors"
	"github.com/psocietyyy/go-gin-kafka/service/product_service/internal/model"
	"github.com/psocietyyy/go-gin-kafka/service/product_service/internal/repository"
)

type TrainSeatService struct {
	trainSeatRepo *repository.TrainSeatRepository
	trainRepo     *repository.TrainRepository
}

func NewTrainSeatService(trainSeatRepo *repository.TrainSeatRepository, trainRepo *repository.TrainRepository) *TrainSeatService {
	return &TrainSeatService{
		trainSeatRepo: trainSeatRepo,
		trainRepo:     trainRepo,
	}
}

func (s *TrainSeatService) FindByID(id uint) (model.TrainSeat, error) {
	return s.trainSeatRepo.FindByID(id)
}

func (s *TrainSeatService) FindByTrainID(trainID uint) []model.TrainSeat {
	return s.trainSeatRepo.FindByTrainID(trainID)
}

func (s *TrainSeatService) Create(seat model.TrainSeat) (model.TrainSeat, error) {
	// Verify if the train exists first
	_, err := s.trainRepo.FindByID(seat.TrainID)
	if err != nil {
		return model.TrainSeat{}, trainError.ErrTrainNotFound
	}
	return s.trainSeatRepo.Create(seat)
}

func (s *TrainSeatService) Update(seat model.TrainSeat) (model.TrainSeat, error) {
	_, err := s.trainSeatRepo.FindByID(seat.ID)
	if err != nil {
		return model.TrainSeat{}, err
	}
	// Verify if the updated train exists
	_, err = s.trainRepo.FindByID(seat.TrainID)
	if err != nil {
		return model.TrainSeat{}, trainError.ErrTrainNotFound
	}
	return s.trainSeatRepo.Update(seat)
}

func (s *TrainSeatService) Delete(id uint) error {
	_, err := s.trainSeatRepo.FindByID(id)
	if err != nil {
		return err
	}
	return s.trainSeatRepo.Delete(id)
}
