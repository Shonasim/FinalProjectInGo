package handlers

import (
	"FinalProject/internal/models"
	errors2 "FinalProject/pkg/errors"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

func (h *Handler) GetRoutes(c *gin.Context) {
	routes, err := h.service.GetRoutes()
	if err != nil {
		h.logger.Printf("GetRoutes - h.service.GetRoutes error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": routes})
}

func (h *Handler) GetRouteById(c *gin.Context) {
	routeIdStr := c.Param("route_id")
	routeId, err := strconv.Atoi(routeIdStr)
	if err != nil {
		h.logger.Printf("GetRouteById - strconv.Atoi error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid route_id"})
		return
	}
	routes, err := h.service.GetRouteById(routeId)
	if err != nil {
		h.logger.Printf("GetRouteById - h.service.GetRouteById error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": routes})
}

func (h *Handler) FinishRoute(c *gin.Context) {
	var finishReq models.Finish
	err := c.BindJSON(&finishReq)
	if err != nil {
		h.logger.Printf("FinishRoute - c.BindJSON error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": errors2.ErrBindJSON})
		return
	}

	err = h.service.FinishRoute(finishReq)
	if err != nil {
		h.logger.Printf("FinishRoute - h.service.FinishRoute error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "internal_error"})
		return
	}

	c.JSON(http.StatusOK, "Successfully finished route")
}
