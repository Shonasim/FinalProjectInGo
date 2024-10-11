package handlers

import (
	"FinalProject/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Register(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		h.logger.Error("Failed to read request: ", err)
		return
	}
	h.service.Registration(user)
}
