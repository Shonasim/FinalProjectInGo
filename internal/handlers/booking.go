package handlers

import (
	"FinalProject/internal/models"
	errors2 "FinalProject/pkg/errors"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) CreateReservation(c *gin.Context) {
	var booking models.Booking
	err := c.BindJSON(&booking)
	if err != nil {
		h.logger.Printf("CreateReservation - c.BindJSON error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": errors2.ErrBindJSON})
		return
	}
	newBooking, err := h.service.AddBooking(&booking)
	if err != nil {
		if errors.Is(err, errors2.ErrInvalidUserID) ||
			errors.Is(err, errors2.ErrInvalidDriverID) ||
			errors.Is(err, errors2.ErrInvalidSeatID) ||
			errors.Is(err, errors2.ErrInvalidStatusID) ||
			errors.Is(err, errors2.ErrInvalidCity) ||
			errors.Is(err, errors2.ErrInvalidPrice) {
			h.logger.Printf("CreateReservation - h.service.AddBooking error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		h.logger.Printf("CreateReservation - h.service.AddBooking error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": newBooking})
}
