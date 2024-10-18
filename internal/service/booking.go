package service

import (
	"FinalProject/internal/models"
	errors2 "FinalProject/pkg/errors"
)

func (s *Service) AddBooking(u *models.Booking) (*models.Booking, error) {
	err := validateBooking(u)
	if err != nil {
		return nil, err
	}
	return s.Repository.AddBooking(u)
}

func (s *Service) GetReservationById(id, userId int) (*models.Booking, error) {
	return s.Repository.GetReservationById(id, userId)
}

func validateBooking(u *models.Booking) error {
	// Проверяем, что все обязательные идентификаторы установлены
	if u.UserId <= 0 {
		return errors2.ErrInvalidUserID
	}
	if u.DriverId <= 0 {
		return errors2.ErrInvalidDriverID
	}
	if u.SeatsId <= 0 {
		return errors2.ErrInvalidSeatID
	}
	if u.StatusId <= 0 {
		return errors2.ErrInvalidStatusID
	}
	if u.StartCityId <= 0 {
		return errors2.ErrInvalidCity
	}
	if u.EndCityId <= 0 {
		return errors2.ErrInvalidCity
	}

	// Проверяем, что цена положительная
	if u.Price <= 0 {
		return errors2.ErrInvalidPrice
	}
	return nil
}

func (s *Service) GetReservation(id int) {

}
