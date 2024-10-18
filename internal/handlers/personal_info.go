package handlers

import (
	"FinalProject/internal/models"
	errors2 "FinalProject/pkg/errors"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) AddPersonalInfo(c *gin.Context) {
	var info models.PersonalInformation
	err := c.BindJSON(&info)
	if err != nil {
		h.logger.Printf("AddPersonalInfo - c.BindJSON error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": errors2.ErrBindJSON})
		return
	}
	newInfo, err := h.service.FillExtraInfo(info)
	if err != nil {
		if errors.Is(err, errors2.ErrInvalidFirstName) || errors.Is(err, errors2.ErrInvalidLastName) || errors.Is(err, errors2.ErrInvalidFathersName) || errors.Is(err, errors2.ErrInvalidAboutUser) || errors.Is(err, errors2.ErrInvalidSex) {
			h.logger.Printf("AddPersonalInfo - h.service.FillExtraInfo error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		h.logger.Printf("AddPersonalInfo - h.service.FillExtraInfo error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": newInfo})
}

func (h *Handler) GetPersonalInfoByID(c *gin.Context) {
	userIDStr := c.Param("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		h.logger.Printf("GetPersonalInfoByID - strconv.Atoi error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": errors2.ErrFailedConvert})
		return
	}

	persInfo, err := h.service.GetPersInfoById(userID)
	if err != nil {
		log.Println("GetPersonalInfoByID - h.service.GetPersInfoById err:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "internal error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": persInfo})
}
