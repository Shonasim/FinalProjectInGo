package handlers

import (
	"FinalProject/internal/models"
	errors2 "FinalProject/pkg/errors"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) AddSeats(c *gin.Context) {
	var seats models.Seats
	if err := c.BindJSON(&seats); err != nil {
		h.logger.Printf("AddSeats - c.BindJSON error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	userIDAny, _ := c.Get("user_id")
	fmt.Println(userIDAny)
	userID := userIDAny.(int)
	seats.CarId = userID

	newSeats, err := h.service.FillSeats(seats)
	if err != nil {
		if errors.Is(err, errors2.ErrInvalidSeats) {
			h.logger.Printf("AddSeats - h.service.FillSeats: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		h.logger.Printf("AddSeats - h.service.FillSeats error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": newSeats})

}
