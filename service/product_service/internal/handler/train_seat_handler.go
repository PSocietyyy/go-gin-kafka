package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	dto "github.com/psocietyyy/go-gin-kafka/service/product_service/internal/dto/train_seat"
	trainError "github.com/psocietyyy/go-gin-kafka/service/product_service/internal/errors"
	"github.com/psocietyyy/go-gin-kafka/service/product_service/internal/model"
	"github.com/psocietyyy/go-gin-kafka/service/product_service/internal/service"
)

type TrainSeatHandler struct {
	trainSeatService *service.TrainSeatService
}

func NewTrainSeatHandler(trainSeatService *service.TrainSeatService) *TrainSeatHandler {
	return &TrainSeatHandler{trainSeatService: trainSeatService}
}

// Get TrainSeat By ID
func (h *TrainSeatHandler) FindByID(c *gin.Context) {
	id := c.Param("id")
	uintId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid ID"})
		return
	}
	seat, err := h.trainSeatService.FindByID(uint(uintId))
	if err != nil {
		if err == trainError.ErrTrainSeatNotFound {
			c.JSON(404, gin.H{"message": "Train Seat Not Found"})
			return
		}
		c.JSON(500, gin.H{"message": "Failed to Get Train Seat By ID"})
		return
	}
	c.JSON(200, gin.H{
		"message": "Get Train Seat By ID Successfully",
		"data":    dto.ToTrainSeatResponse(seat),
	})
}

// Get TrainSeats By Train ID
func (h *TrainSeatHandler) FindByTrainID(c *gin.Context) {
	trainID := c.Param("id")
	uintId, err := strconv.ParseUint(trainID, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid Train ID"})
		return
	}
	seats := h.trainSeatService.FindByTrainID(uint(uintId))
	c.JSON(200, gin.H{
		"message": "Get Train Seats By Train ID Successfully",
		"data":    dto.ToTrainSeatResponses(seats),
	})
}

// Create TrainSeat
func (h *TrainSeatHandler) Create(c *gin.Context) {
	var request dto.TrainSeatCreateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"message": "Invalid Request"})
		return
	}
	seat := model.TrainSeat{
		TrainID:    request.TrainID,
		SeatNumber: request.SeatNumber,
	}
	seat, err := h.trainSeatService.Create(seat)
	if err != nil {
		if err == trainError.ErrTrainNotFound {
			c.JSON(404, gin.H{"message": "Train Not Found"})
			return
		}
		c.JSON(500, gin.H{"message": "Failed to Create Train Seat", "error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"message": "Train Seat Created Successfully",
		"data":    dto.ToTrainSeatResponse(seat),
	})
}

// Update TrainSeat
func (h *TrainSeatHandler) Update(c *gin.Context) {
	id := c.Param("id")
	uintId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid ID"})
		return
	}
	var request dto.TrainSeatUpdateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"message": "Invalid Request"})
		return
	}
	seat, err := h.trainSeatService.FindByID(uint(uintId))
	if err != nil {
		if err == trainError.ErrTrainSeatNotFound {
			c.JSON(404, gin.H{"message": "Train Seat Not Found"})
			return
		}
		c.JSON(500, gin.H{"message": "Failed to Get Train Seat By ID"})
		return
	}
	
	if request.TrainID != 0 {
		seat.TrainID = request.TrainID
	}
	if request.SeatNumber != "" {
		seat.SeatNumber = request.SeatNumber
	}

	seat, err = h.trainSeatService.Update(seat)
	if err != nil {
		c.JSON(500, gin.H{"message": "Failed to Update Train Seat", "error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"message": "Train Seat Updated Successfully",
		"data":    dto.ToTrainSeatResponse(seat),
	})
}

// Delete TrainSeat
func (h *TrainSeatHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	uintId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid ID"})
		return
	}
	err = h.trainSeatService.Delete(uint(uintId))
	if err != nil {
		if err == trainError.ErrTrainSeatNotFound {
			c.JSON(404, gin.H{"message": "Train Seat Not Found"})
			return
		}
		c.JSON(500, gin.H{"message": "Failed to Delete Train Seat"})
		return
	}
	c.JSON(200, gin.H{
		"message": "Train Seat Deleted Successfully",
	})
}
