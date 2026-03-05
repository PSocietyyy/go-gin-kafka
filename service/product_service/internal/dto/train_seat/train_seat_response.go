package dto

import (
	trainDto "github.com/psocietyyy/go-gin-kafka/service/product_service/internal/dto/train"
	"github.com/psocietyyy/go-gin-kafka/service/product_service/internal/model"
)

type TrainSeatResponse struct {
	ID         uint                   `json:"id"`
	TrainID    uint                   `json:"train_id"`
	SeatNumber string                 `json:"seat_number"`
	Train      trainDto.TrainResponse `json:"train"`
}

func ToTrainSeatResponse(seat model.TrainSeat) TrainSeatResponse {
	return TrainSeatResponse{
		ID:         seat.ID,
		TrainID:    seat.TrainID,
		SeatNumber: seat.SeatNumber,
		Train:      trainDto.ToTrainResponse(seat.Train),
	}
}

func ToTrainSeatResponses(seats []model.TrainSeat) []TrainSeatResponse {
	var responses []TrainSeatResponse
	for _, seat := range seats {
		responses = append(responses, ToTrainSeatResponse(seat))
	}
	if len(responses) == 0 {
		return []TrainSeatResponse{}
	}
	return responses
}
