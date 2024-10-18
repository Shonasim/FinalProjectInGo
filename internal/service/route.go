package service

import (
	"FinalProject/internal/models"
	errors2 "FinalProject/pkg/errors"
)

func (s *Service) AddRoute(route *models.Route) (*models.Route, error) {
	err := validateRoute(route)
	if err != nil {
		return nil, err
	}
	return s.Repository.AddRoute(route)
}

func (s *Service) GetRoutes() ([]models.Route, error) {
	return s.Repository.GetRoutes()
}

func validateRoute(route *models.Route) error {
	if route.FromCity == 0 {
		return errors2.ErrInvalidCity
	}
	if route.ToCity == 0 {
		return errors2.ErrInvalidCity
	}
	if route.Price == 0 {
		return errors2.ErrInvalidPrice
	}
	if route.CarId == 0 {
		return errors2.ErrInvalidCarId
	}
	return nil
}
