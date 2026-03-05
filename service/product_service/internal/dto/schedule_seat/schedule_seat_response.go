package dto

import (
	scheduleDto "github.com/psocietyyy/go-gin-kafka/service/product_service/internal/dto/schedule"
	trainSeatDto "github.com/psocietyyy/go-gin-kafka/service/product_service/internal/dto/train_seat"
	"github.com/psocietyyy/go-gin-kafka/service/product_service/internal/model"
)

type ScheduleSeatResponse struct {
	ID          uint                           `json:"id"`
	ScheduleID  uint                           `json:"schedule_id"`
	TrainSeatID uint                           `json:"train_seat_id"`
	Status      string                         `json:"status"`
	Schedule    scheduleDto.ScheduleResponse   `json:"schedule"`
	TrainSeat   trainSeatDto.TrainSeatResponse `json:"train_seat"`
}

func ToScheduleSeatResponse(seat model.ScheduleSeat) ScheduleSeatResponse {
	return ScheduleSeatResponse{
		ID:          seat.ID,
		ScheduleID:  seat.ScheduleID,
		TrainSeatID: seat.TrainSeatID,
		Status:      string(seat.Status),
		Schedule:    scheduleDto.ToScheduleResponse(seat.Schedule),
		TrainSeat:   trainSeatDto.ToTrainSeatResponse(seat.TrainSeat),
	}
}

func ToScheduleSeatResponses(seats []model.ScheduleSeat) []ScheduleSeatResponse {
	var responses []ScheduleSeatResponse
	for _, seat := range seats {
		responses = append(responses, ToScheduleSeatResponse(seat))
	}
	if len(responses) == 0 {
		return []ScheduleSeatResponse{}
	}
	return responses
}
