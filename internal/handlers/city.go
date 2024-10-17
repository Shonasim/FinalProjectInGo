package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetCities(c *gin.Context) {
	cities, err := h.service.GetServiceList()
	if err != nil {
		h.logger.Printf("GetCities - h.service.GetServiceList error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"cities": cities})
}
