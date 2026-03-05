package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	dto "github.com/psocietyyy/go-gin-kafka/service/product_service/internal/dto/schedule"
	trainError "github.com/psocietyyy/go-gin-kafka/service/product_service/internal/errors"
	"github.com/psocietyyy/go-gin-kafka/service/product_service/internal/model"
	"github.com/psocietyyy/go-gin-kafka/service/product_service/internal/service"
)

type ScheduleHandler struct {
	scheduleService *service.ScheduleService
}

func NewScheduleHandler(scheduleService *service.ScheduleService) *ScheduleHandler {
	return &ScheduleHandler{scheduleService: scheduleService}
}

// Get All Schedules
func (h *ScheduleHandler) FindAll(c *gin.Context) {
	schedules := h.scheduleService.FindAll()
	c.JSON(200, gin.H{
		"message": "Get All Schedules Successfully",
		"data":    dto.ToScheduleResponses(schedules),
	})
}

// Get Schedule By ID
func (h *ScheduleHandler) FindByID(c *gin.Context) {
	id := c.Param("id")
	uintId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid ID"})
		return
	}
	schedule, err := h.scheduleService.FindByID(uint(uintId))
	if err != nil {
		if err == trainError.ErrScheduleNotFound {
			c.JSON(404, gin.H{"message": "Schedule Not Found"})
			return
		}
		c.JSON(500, gin.H{"message": "Failed to Get Schedule By ID"})
		return
	}
	c.JSON(200, gin.H{
		"message": "Get Schedule By ID Successfully",
		"data":    dto.ToScheduleResponse(schedule),
	})
}

// Create Schedule
func (h *ScheduleHandler) Create(c *gin.Context) {
	var request dto.ScheduleCreateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"message": "Invalid Request"})
		return
	}
	schedule := model.Schedule{
		TrainID:       request.TrainID,
		DepartureTime: request.DepartureTime,
		ArrivalTime:   request.ArrivalTime,
	}
	schedule, err := h.scheduleService.Create(schedule)
	if err != nil {
		c.JSON(500, gin.H{"message": "Failed to Create Schedule", "error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"message": "Schedule Created Successfully",
		"data":    dto.ToScheduleResponse(schedule),
	})
}

// Update Schedule
func (h *ScheduleHandler) Update(c *gin.Context) {
	id := c.Param("id")
	uintId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid ID"})
		return
	}
	var request dto.ScheduleUpdateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"message": "Invalid Request"})
		return
	}
	schedule, err := h.scheduleService.FindByID(uint(uintId))
	if err != nil {
		if err == trainError.ErrScheduleNotFound {
			c.JSON(404, gin.H{"message": "Schedule Not Found"})
			return
		}
		c.JSON(500, gin.H{"message": "Failed to Get Schedule By ID"})
		return
	}
	
	if request.TrainID != 0 {
		schedule.TrainID = request.TrainID
	}
	if !request.DepartureTime.IsZero() {
		schedule.DepartureTime = request.DepartureTime
	}
	if !request.ArrivalTime.IsZero() {
		schedule.ArrivalTime = request.ArrivalTime
	}

	schedule, err = h.scheduleService.Update(schedule)
	if err != nil {
		c.JSON(500, gin.H{"message": "Failed to Update Schedule", "error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"message": "Schedule Updated Successfully",
		"data":    dto.ToScheduleResponse(schedule),
	})
}

// Delete Schedule
func (h *ScheduleHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	uintId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid ID"})
		return
	}
	err = h.scheduleService.Delete(uint(uintId))
	if err != nil {
		if err == trainError.ErrScheduleNotFound {
			c.JSON(404, gin.H{"message": "Schedule Not Found"})
			return
		}
		c.JSON(500, gin.H{"message": "Failed to Delete Schedule"})
		return
	}
	c.JSON(200, gin.H{
		"message": "Schedule Deleted Successfully",
	})
}
