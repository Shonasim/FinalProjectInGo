package service

import (
	"FinalProject/internal/models"
	errors2 "FinalProject/pkg/errors"
	"time"
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

func (s *Service) GetRouteById(id int) (*models.Route, error) {
	return s.Repository.GetRouteByID(id)
}

func (s *Service) FinishRoute(u models.Finish) error {
	return s.Repository.FinishRoute(u)
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
	layout := "2006-01-02"
	d1, err := time.Parse(layout, time.Now().String())
	if err != nil {
		return errors2.ErrInvalidDate
	}
	d2, err := time.Parse(layout, route.Date.String())
	if err != nil {
		return errors2.ErrInvalidDate
	}
	if d1.After(d2) {
		return errors2.ErrInvalidDate
	}
	return nil
}
