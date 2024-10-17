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
		h.logger.Printf("AddCar - c.BindJSON error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	// Заполнение или обработка данных о машине
	newCar, err := h.service.FillCars(car)
	if err != nil {
		if errors.Is(err, errors2.ErrInvalidModel) ||
			errors.Is(err, errors2.ErrInvalidMark) ||
			errors.Is(err, errors2.ErrInvalidAutobody) ||
			errors.Is(err, errors2.ErrInvalidSeats) ||
			errors.Is(err, errors2.ErrInvalidCarNumber) ||
			errors.Is(err, errors2.ErrInvalidAutobody) {
			h.logger.Printf("AddCar - h.service.AddCar error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		h.logger.Printf("AddCar - h.service.AddCar error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": newCar})

}
