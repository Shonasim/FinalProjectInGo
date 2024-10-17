package handlers

import (
	"FinalProject/internal/models"
	errors2 "FinalProject/pkg/errors"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) CreateRoute(c *gin.Context) {
	var route models.Route
	err := c.BindJSON(&route)
	if err != nil {
		h.logger.Printf("CreateRoute - c.BindJSON error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": errors2.ErrBindJSON})
		return
	}
	userIDAny, _ := c.Get("user_id")
	userID := userIDAny.(int)
	route.DriverId = userID
	addRoute, err := h.service.AddRoute(&route)
	if err != nil {
		if errors.Is(err, errors2.ErrInvalidCity) || errors.Is(err, errors2.ErrInvalidPrice) || errors.Is(err, errors2.ErrInvalidCarId) {
			h.logger.Printf("CreateRoute - h.service.AddRoute error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		h.logger.Printf("CreateRoute - h.service.AddRoute error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": addRoute})
}
