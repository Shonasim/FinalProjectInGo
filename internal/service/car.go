package service

import (
	"FinalProject/internal/models"
	errors2 "FinalProject/pkg/errors"
)

func (s *Service) FillCars(car models.Car) (*models.Car, error) {
	err := validateCar(&car)
	if err != nil {
		return nil, err
	}
	return s.Repository.AddCar(&car)
}
func validateCar(car *models.Car) error {
	if car.Model == "" {
		return errors2.ErrInvalidModel
	}
	if car.Mark == "" {
		return errors2.ErrInvalidMark
	}
	if car.Autobody == "" {
		return errors2.ErrInvalidAutobody
	}
	x := len(car.CarNumber)
	if (car.CarNumber == "") && (x <= 6) {
		return errors2.ErrInvalidCarNumber
	}
	return nil
}
