package dto

import (
	"time"

	trainDto "github.com/psocietyyy/go-gin-kafka/service/product_service/internal/dto/train"
	"github.com/psocietyyy/go-gin-kafka/service/product_service/internal/model"
)

type ScheduleResponse struct {
	ID            uint                   `json:"id"`
	TrainID       uint                   `json:"train_id"`
	Train         trainDto.TrainResponse `json:"train"`
	DepartureTime time.Time              `json:"departure_time"`
	ArrivalTime   time.Time              `json:"arrival_time"`
}

func ToScheduleResponse(schedule model.Schedule) ScheduleResponse {
	return ScheduleResponse{
		ID:            schedule.ID,
		TrainID:       schedule.TrainID,
		Train:         trainDto.ToTrainResponse(schedule.Train),
		DepartureTime: schedule.DepartureTime,
		ArrivalTime:   schedule.ArrivalTime,
	}
}

func ToScheduleResponses(schedules []model.Schedule) []ScheduleResponse {
	var responses []ScheduleResponse
	for _, schedule := range schedules {
		responses = append(responses, ToScheduleResponse(schedule))
	}
	if len(responses) == 0 {
		return []ScheduleResponse{}
	}
	return responses
}
