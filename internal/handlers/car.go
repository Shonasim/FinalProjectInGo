package handlers

import (
	"FinalProject/internal/models"
	errors2 "FinalProject/pkg/errors"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) CreateCar(c *gin.Context) {
	var car models.Car
	err := c.BindJSON(&car)
	if err != nil {
		h.logger.Printf("CreateCar - c.BindJSON error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": errors2.ErrBindJSON})
		return
	}

	userIDAny, _ := c.Get("user_id")
	userID := userIDAny.(int)
	car.UserId = userID

	// Заполнение или обработка данных о машине
	newCar, err := h.service.FillCars(car)
	if err != nil {
		if errors.Is(err, errors2.ErrInvalidModel) ||
			errors.Is(err, errors2.ErrInvalidMark) ||
			errors.Is(err, errors2.ErrInvalidAutobody) ||
			errors.Is(err, errors2.ErrInvalidSeats) ||
			errors.Is(err, errors2.ErrInvalidCarNumber) ||
			errors.Is(err, errors2.ErrInvalidAutobody) {
			h.logger.Printf("CreateCar - h.service.FillCars error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		h.logger.Printf("CreateCar - h.service.AddCar error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": newCar})
}

func (h *Handler) GetCars(c *gin.Context) {
	userIDAny, _ := c.Get("user_id")
	userID := userIDAny.(int)
	// Заполнение или обработка данных о машине
	newCar, err := h.service.GetCars(userID)
	if err != nil {
		h.logger.Printf("GetCars - h.service.GetCars error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": newCar})
}
