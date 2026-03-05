package dto

import "time"

type ScheduleCreateRequest struct {
	TrainID       uint      `json:"train_id" binding:"required"`
	DepartureTime time.Time `json:"departure_time" binding:"required"`
	ArrivalTime   time.Time `json:"arrival_time" binding:"required"`
}

type ScheduleUpdateRequest struct {
	TrainID       uint      `json:"train_id" binding:"omitempty"`
	DepartureTime time.Time `json:"departure_time" binding:"omitempty"`
	ArrivalTime   time.Time `json:"arrival_time" binding:"omitempty"`
}
