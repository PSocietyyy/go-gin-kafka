package dto

import "github.com/psocietyyy/go-gin-kafka/service/order_service/internal/model"

type OrderResponse struct {
	ID uint `json:"id"`
	UserID uint `json:"user_id"`
	ScheduleSeatID uint `json:"schedule_seat_id"`
	Status string `json:"status"`
}

func ToOrderResponse(order *model.Order) OrderResponse {
	return OrderResponse{
		ID: order.ID,
		UserID: order.UserID,
		ScheduleSeatID: order.ScheduleSeatID,
		Status: string(order.Status),
	}
}

func ToOrderResponses(orders []model.Order) []OrderResponse {
	var responses []OrderResponse
	for _, order := range orders {
		responses = append(responses, ToOrderResponse(&order))
	}
	if len(responses) == 0 {
		return []OrderResponse{}
	}
	return responses
}