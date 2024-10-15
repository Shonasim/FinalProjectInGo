package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *Handler) GetPersonalInfoUsers(c *gin.Context) {
	users, err := h.service.ListUsers()
	if err != nil {
		log.Printf("GetUsers - h.service.ListUsers error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}
func (h *Handler) GetUserByLastName() {

}
