package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	dto "github.com/psocietyyy/go-gin-kafka/service/product_service/internal/dto/schedule_seat"
	trainError "github.com/psocietyyy/go-gin-kafka/service/product_service/internal/errors"
	"github.com/psocietyyy/go-gin-kafka/service/product_service/internal/model"
	"github.com/psocietyyy/go-gin-kafka/service/product_service/internal/service"
)

type ScheduleSeatHandler struct {
	scheduleSeatService *service.ScheduleSeatService
}

func NewScheduleSeatHandler(scheduleSeatService *service.ScheduleSeatService) *ScheduleSeatHandler {
	return &ScheduleSeatHandler{scheduleSeatService: scheduleSeatService}
}

// Get All Schedule Seats
func (h *ScheduleSeatHandler) FindAll(c *gin.Context) {
	seats := h.scheduleSeatService.FindAll()
	c.JSON(200, gin.H{
		"message": "Get All Schedule Seats Successfully",
		"data":    dto.ToScheduleSeatResponses(seats),
	})
}

// Get ScheduleSeat By ID
func (h *ScheduleSeatHandler) FindByID(c *gin.Context) {
	id := c.Param("id")
	uintId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid ID"})
		return
	}
	seat, err := h.scheduleSeatService.FindByID(uint(uintId))
	if err != nil {
		if err == trainError.ErrScheduleSeatNotFound {
			c.JSON(404, gin.H{"message": "Schedule Seat Not Found"})
			return
		}
		c.JSON(500, gin.H{"message": "Failed to Get Schedule Seat By ID"})
		return
	}
	c.JSON(200, gin.H{
		"message": "Get Schedule Seat By ID Successfully",
		"data":    dto.ToScheduleSeatResponse(seat),
	})
}

// Get ScheduleSeats By Schedule ID
func (h *ScheduleSeatHandler) FindByScheduleID(c *gin.Context) {
	scheduleID := c.Param("id")
	uintId, err := strconv.ParseUint(scheduleID, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid Schedule ID"})
		return
	}
	seats := h.scheduleSeatService.FindByScheduleID(uint(uintId))
	c.JSON(200, gin.H{
		"message": "Get Schedule Seats By Schedule ID Successfully",
		"data":    dto.ToScheduleSeatResponses(seats),
	})
}

// Create ScheduleSeat
func (h *ScheduleSeatHandler) Create(c *gin.Context) {
	var request dto.ScheduleSeatCreateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"message": "Invalid Request"})
		return
	}
	
	// Default status is available if not specified
	status := request.Status
	if status == "" {
		status = model.SeatStatusAvailable
	}
	
	seat := model.ScheduleSeat{
		ScheduleID:  request.ScheduleID,
		TrainSeatID: request.TrainSeatID,
		Status:      status,
	}
	seat, err := h.scheduleSeatService.Create(seat)
	if err != nil {
		c.JSON(500, gin.H{"message": "Failed to Create Schedule Seat", "error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"message": "Schedule Seat Created Successfully",
		"data":    dto.ToScheduleSeatResponse(seat),
	})
}

// Book a Seat
func (h *ScheduleSeatHandler) BookSeat(c *gin.Context) {
	id := c.Param("id")
	uintId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid ID"})
		return
	}
	
	err = h.scheduleSeatService.BookSeat(uint(uintId))
	if err != nil {
		if err == trainError.ErrScheduleSeatNotFound {
			c.JSON(404, gin.H{"message": "Schedule Seat Not Found"})
			return
		}
		if err == trainError.ErrSeatAlreadyBooked {
			c.JSON(409, gin.H{"message": "Seat Already Booked"})
			return
		}
		c.JSON(500, gin.H{"message": "Failed to Book Seat", "error": err.Error()})
		return
	}
	
	c.JSON(200, gin.H{
		"message": "Seat Booked Successfully",
	})
}

// Cancel Booking
func (h *ScheduleSeatHandler) CancelSeat(c *gin.Context) {
	id := c.Param("id")
	uintId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid ID"})
		return
	}
	
	err = h.scheduleSeatService.CancelSeat(uint(uintId))
	if err != nil {
		if err == trainError.ErrScheduleSeatNotFound {
			c.JSON(404, gin.H{"message": "Schedule Seat Not Found"})
			return
		}
		c.JSON(500, gin.H{"message": "Failed to Cancel Seat Booking", "error": err.Error()})
		return
	}
	
	c.JSON(200, gin.H{
		"message": "Seat Booking Cancelled Successfully",
	})
}
