package dto

type OrderCreateRequest struct {
	UserID uint `json:"user_id" binding:"required"`
	ScheduleSeatID uint `json:"schedule_seat_id" binding:"required"`
}

type OrderUpdateRequest struct {
	Status string `json:"status" binding:"required"`
}