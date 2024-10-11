package service

import (
	"FinalProject/internal/models"
	"FinalProject/internal/utils"
	"FinalProject/pkg/errors"
)

func (s Service) Registration(user models.User) (models.User, error) {
	err := validateUser(user)
	if err != nil {
		return models.User{}, err
	}
	hashPass, err := utils.HashPassword(user.Password)
	if err != nil {
		return models.User{}, errors.ErrFailedHashing
	}
	user.Password = hashPass
	s.Repository.AddUser(&user)
}

func validateUser(user models.User) error {
	if user.FirstName == "" {
		return errors.ErrInvalidFirstName
	}
	if user.LastName == "" {
		return errors.ErrInvalidSecondName
	}
	if user.FathersName == "" {
		return false
	}
	if user.Email == "" {
		return false
	}
	if user.Password == "" {
		return false
	}
	return true
}
