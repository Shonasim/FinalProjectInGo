package repository

import (
	"FinalProject/internal/models"
	"FinalProject/pkg/errors"
)

func (r *Repository) AddBooking(u *models.Booking) (*models.Booking, error) {
	query := `update seats set available = false where seat_id = ?`
	err := r.db.Exec(query, u.SeatsId).Error
	if err != nil {
		r.logger.Error("Faced an error while tried to update seatings, err: ", err)
		return nil, errors.ErrFailedToInsert
	}

	query = `INSERT INTO bookings
    (user_id, driver_id,car_id, route_id, seats_id, status_id, price, start_city_id, end_city_id)
	VALUES (?,?,?,?,?,?,?,?)`
	err = r.db.Exec(query, u.UserId, u.DriverId, u.CarId, u.RouteId, u.SeatsId, u.StatusId, u.Price, u.StartCityId, u.EndCityId).Error
	if err != nil {
		r.logger.Error("Faced an error while tried to insert booking, err: ", err)
		return nil, errors.ErrFailedToInsert
	}
	return u, nil
}

func (r *Repository) GetReservationById(id, userId int) (*models.Booking, error) {
	var reservation models.Booking
	query := `select * from bookings where booking_id = ? and user_id = ?`
	err := r.db.Raw(query, id, userId).Scan(&reservation).Error
	if err != nil {
		r.logger.Error("Faced an error while tried to select booking, err: ", err)
		return nil, errors.ErrFailedToGet
	}
	return &reservation, nil
}
