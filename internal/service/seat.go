package service

import (
	"FinalProject/internal/models"
	errors2 "FinalProject/pkg/errors"
)

func (s *Service) FillSeats(seats models.Seats) (*models.Seats, error) {
	err := validateSeat(&seats)
	if err != nil {
		return nil, err
	}
	return s.Repository.AddSeat(&seats)
}

func validateSeat(s *models.Seats) error {
	x := len(s.Seats)
	if x == 0 {
		return errors2.ErrInvalidSeats
	}
	return nil
}
