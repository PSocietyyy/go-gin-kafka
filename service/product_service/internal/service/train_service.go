package service

import (
	trainError "github.com/psocietyyy/go-gin-kafka/service/product_service/internal/errors"
	"github.com/psocietyyy/go-gin-kafka/service/product_service/internal/model"
	"github.com/psocietyyy/go-gin-kafka/service/product_service/internal/repository"
)

type TrainService struct {
	trainRepo *repository.TrainRepository
}

func NewTrainService(trainRepo *repository.TrainRepository) *TrainService {
	return &TrainService{trainRepo: trainRepo}
}

func (s *TrainService) FindAll() []model.Train {
	return s.trainRepo.FindAll()
}

func (s *TrainService) FindByID(id uint) (model.Train, error) {
	return s.trainRepo.FindByID(id)
}

func (s *TrainService) Create(train model.Train) (model.Train, error) {
	_, err := s.trainRepo.FindByCode(train.Code)
	if err == nil {
		return model.Train{}, trainError.ErrTrainCodeAlreadyExists
	}
	return s.trainRepo.Create(train)
}

func (s *TrainService) Update(train model.Train) (model.Train, error) {
	_, err := s.trainRepo.FindByID(train.ID)
	if err != nil {
		return model.Train{}, err
	}
	return s.trainRepo.Update(train)
}

func (s *TrainService) Delete(id uint) error {
	_, err := s.trainRepo.FindByID(id)
	if err != nil {
		return err
	}
	return s.trainRepo.Delete(id)
}

func (s *TrainService) FindByCode(code string) (model.Train, error) {
	return s.trainRepo.FindByCode(code)
}