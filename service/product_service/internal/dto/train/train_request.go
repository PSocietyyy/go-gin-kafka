package dto

type TrainCreateRequest struct {
	Name	string	`json:"name" binding:"required"`
	Code	string	`json:"code" binding:"required"`
}

type TrainUpdateRequest struct {
	Name	string	`json:"name" binding:"omitempty"`
	Code	string	`json:"code" binding:"omitempty"`
}