package dto

type TrainSeatCreateRequest struct {
	TrainID    uint   `json:"train_id" binding:"required"`
	SeatNumber string `json:"seat_number" binding:"required"`
}

type TrainSeatUpdateRequest struct {
	TrainID    uint   `json:"train_id" binding:"omitempty"`
	SeatNumber string `json:"seat_number" binding:"omitempty"`
}
