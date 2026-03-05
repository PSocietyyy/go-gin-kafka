package dto

import "github.com/psocietyyy/go-gin-kafka/service/product_service/internal/model"

type ScheduleSeatCreateRequest struct {
	ScheduleID  uint             `json:"schedule_id" binding:"required"`
	TrainSeatID uint             `json:"train_seat_id" binding:"required"`
	Status      model.SeatStatus `json:"status" binding:"omitempty"`
}
