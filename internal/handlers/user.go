package handlers

import (
	"FinalProject/internal/models"
	errors2 "FinalProject/pkg/errors"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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

func (h *Handler) SignIn(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		h.logger.Printf("SignIn - c.BindJSON error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": errors2.ErrBindJSON})
		return
	}

	//log.Printf("SignIn - data after binding: %v", user)

	token, err := h.service.SignIn(&user)
	if err != nil {
		h.logger.Printf("SignIn - h.service.SignIn error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
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
