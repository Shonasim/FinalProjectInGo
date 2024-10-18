package handlers

import (
	"FinalProject/internal/models"
	errors2 "FinalProject/pkg/errors"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

const uploadPath = "./uploads"

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
		h.logger.Println("GetPersonalInfoByID - h.service.GetPersInfoById err:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "internal error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": persInfo})
}

func (h *Handler) UploadPhoto(c *gin.Context) {
	userIDAny, _ := c.Get("user_id")
	userID := userIDAny.(int)

	file, err := c.FormFile("file")
	if err != nil {
		h.logger.Printf("UploadPhoto - c.FormFile error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to get the file"})
		return
	}

	filename := filepath.Join(uploadPath, fmt.Sprintf("%s.jpg", string(userID)))

	if err := c.SaveUploadedFile(file, filename); err != nil {
		h.logger.Printf("UploadPhoto - c.SaveUploadedFile error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save"})
		return
	}

	err = h.service.UploadPhoto(userID, filename)
	if err != nil {
		h.logger.Printf("UploadPhoto - h.service.UploadPhoto error: %v", err)
		c.JSON(http.StatusCreated, gin.H{"warn": "saved photo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Successfully saved"})
}

func (h *Handler) GetPhoto(c *gin.Context) {
	userID := c.Param("user_id")

	// Путь к файлу
	filename := filepath.Join(uploadPath, fmt.Sprintf("%s.jpg", userID))

	// Проверяем, существует ли файл
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		h.logger.Printf("GetPhoto - os.Stat error: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "File does not exists"})
		return
	}

	// Открываем файл
	file, err := os.Open(filename)
	if err != nil {
		h.logger.Printf("GetPhoto - os.Open error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to open the file"})
		return
	}
	defer file.Close()

	// Отправляем файл на фронт
	c.Writer.Header().Set("Content-Type", "image/jpeg")
	c.Writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s.jpg", userID))
	io.Copy(c.Writer, file)
}
