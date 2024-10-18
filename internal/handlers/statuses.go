package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetStatuses(c *gin.Context) {
	statuses, err := h.service.GetStatuses()
	if err != nil {
		h.logger.Printf("GetStatuses - h.service.GetStatuses error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": statuses})
}
