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

func (h *Handler) Register(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		log.Printf("Register - c.BindJSON error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": errors2.ErrBindJSON})
		return
	}
	newUser, err := h.service.Registration(user)
	if err != nil {
		if errors.Is(err, errors2.ErrInvalidEmail) || errors.Is(err, errors2.ErrInvalidPassword) {
			log.Printf("Register - h.service.Registration error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		log.Printf("Register - h.service.Registration error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": models.UserCreateResponse{Email: newUser.Email}})
}

func (h *Handler) AddPersonalInfo(c *gin.Context) {
	var info models.PersonalInformation
	err := c.BindJSON(&info)
	if err != nil {
		log.Printf("AddPersonalInfo - c.BindJSON error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": errors2.ErrBindJSON})
		return
	}
	newInfo, err := h.service.FillExtraInfo(info)
	if err != nil {
		if errors.Is(err, errors2.ErrInvalidFirstName) || errors.Is(err, errors2.ErrInvalidLastName) || errors.Is() {
			log.Printf("Register - h.service.Registration error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		log.Printf("Register - h.service.Registration error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": newInfo})
}

func (h *Handler) GetUsers(c *gin.Context) {
	users, err := h.service.ListUsers()
	if err != nil {
		log.Printf("GetUsers - h.service.ListUsers error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}

func (h *Handler) GetUserByID(c *gin.Context) {
	idStr := c.Param("id")

	if idStr == "" {
		log.Printf("GetUserByID - id is required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("GetUserByID - strconv.Atoi error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.FindUser(id)
	if err != nil {
		log.Printf("GetUserByID - h.service.GetUserByID error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (h *Handler) UpdateUser(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		log.Printf("UpdateUser - c.BindJSON error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateUser, err := h.service.EditUser(&user)
	if err != nil {
		log.Printf("UpdateUser - h.service.UpdateUser error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": updateUser})
}

func (h *Handler) DeleteUser(c *gin.Context) {
	idStr := c.Param("id")

	if idStr == "" {
		log.Printf("DeleteUser - id is required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("DeleteUser - strconv.Atoi error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := h.service.DeleteUser(id); err != nil {
		log.Printf("DeleteUser - h.service.DeleteUser error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}

func (h *Handler) GetUserByEmail(c *gin.Context) {
	email := c.Query("email")

	if email == "" {
		log.Printf("GetUserByEmail - email is required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "email is required"})
		return
	}

	if err := h.service.GetUserByEmail(email); err != nil {
		if errors.Is(err, errors2.ErrInvalidEmail) || errors.Is(err, errors2.ErrAlreadyExists) {
			log.Printf("GetUserByEmail - h.service.GetUserByEmail: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		log.Printf("GetUserByEmail - h.service.GetUserByEmail: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User can create an account"})
}
