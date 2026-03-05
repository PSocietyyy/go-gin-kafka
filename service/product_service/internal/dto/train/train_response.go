package dto

import "github.com/psocietyyy/go-gin-kafka/service/product_service/internal/model"

type TrainResponse struct {
	ID		uint	`json:"id"`
	Name	string	`json:"name"`
	Code	string	`json:"code"`
}

func ToTrainResponse(train model.Train) TrainResponse {
	return TrainResponse{
		ID: train.ID,
		Name: train.Name,
		Code: train.Code,
	}
}

func ToTrainResponses(trains []model.Train) []TrainResponse {
	var trainResponses []TrainResponse
	for _, train := range trains {
		trainResponses = append(trainResponses, ToTrainResponse(train))
	}
	return trainResponses
}