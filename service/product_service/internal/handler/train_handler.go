package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	dto "github.com/psocietyyy/go-gin-kafka/service/product_service/internal/dto/train"
	trainError "github.com/psocietyyy/go-gin-kafka/service/product_service/internal/errors"
	"github.com/psocietyyy/go-gin-kafka/service/product_service/internal/model"
	"github.com/psocietyyy/go-gin-kafka/service/product_service/internal/service"
)

type TrainHandler struct {
	trainService *service.TrainService
}

func NewTrainHandler(trainService *service.TrainService) *TrainHandler {
	return &TrainHandler{trainService: trainService}
}

// Get All Trains
func (h *TrainHandler) FindAll(c *gin.Context) {
	trains := h.trainService.FindAll()
	c.JSON(200, gin.H{
		"message": "Get All Trains Successfully",
		"data": dto.ToTrainResponses(trains),
	})
}

// Get Train By ID
func (h *TrainHandler) FindByID(c *gin.Context) {
	id := c.Param("id")
	uintId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid ID",
		})
		return
	}
	train, err := h.trainService.FindByID(uint(uintId))
	if err != nil {
		if err == trainError.ErrTrainNotFound {
			c.JSON(404, gin.H{
				"message": "Train Not Found",
			})
			return
		}
		c.JSON(500, gin.H{
			"message": "Failed to Get Train By ID",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Get Train By ID Successfully",
		"data": dto.ToTrainResponse(train),
	})
}

// Create Train
func (h *TrainHandler) Create(c *gin.Context) {
	var trainRequest dto.TrainCreateRequest
	if err := c.ShouldBindJSON(&trainRequest); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid Request",
		})
		return
	}
	train := model.Train{
		Name: trainRequest.Name,
		Code: trainRequest.Code,
	}
	train, err := h.trainService.Create(train)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Failed to Create Train",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Train Created Successfully",
		"data": dto.ToTrainResponse(train),
	})
}

// Update Train
func (h *TrainHandler) Update(c *gin.Context) {
	id := c.Param("id")
	uintId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid ID",
		})
		return
	}
	var trainRequest dto.TrainUpdateRequest
	if err := c.ShouldBindJSON(&trainRequest); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid Request",
		})
		return
	}
	train, err := h.trainService.FindByID(uint(uintId))
	if err != nil {
		if err == trainError.ErrTrainNotFound {
			c.JSON(404, gin.H{
				"message": "Train Not Found",
			})
			return
		}
		c.JSON(500, gin.H{
			"message": "Failed to Get Train By ID",
		})
		return
	}
	train.Name = trainRequest.Name
	train.Code = trainRequest.Code
	train, err = h.trainService.Update(train)
	if err != nil {
		if err == trainError.ErrTrainCodeAlreadyExists {
			c.JSON(400, gin.H{
				"message": "Train Code Already Exists",
			})
			return
		}
		c.JSON(500, gin.H{
			"message": "Failed to Update Train",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Train Updated Successfully",
		"data": dto.ToTrainResponse(train),
	})
}

// Delete Train
func (h *TrainHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	uintId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid ID",
		})
		return
	}
	err = h.trainService.Delete(uint(uintId))
	if err != nil {
		if err == trainError.ErrTrainNotFound {
			c.JSON(404, gin.H{
				"message": "Train Not Found",
			})
			return
		}
		c.JSON(500, gin.H{
			"message": "Failed to Delete Train",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Train Deleted Successfully",
	})
}