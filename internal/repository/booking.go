package repository

import (
	"FinalProject/internal/models"
	"FinalProject/pkg/errors"
)

func (r *Repository) AddBooking(u *models.Booking) (*models.Booking, error) {
	query := `INSERT INTO bookings
    (user_id, driver_id, seats_id, status_id, price, start_city_id, end_city_id)
	VALUES (?,?,?,?,?,?,?)`
	err := r.db.Exec(query, u.UserId, u.DriverId, u.SeatsId, u.StatusId, u.Price, u.StartCityId, u.EndCityId).Error
	if err != nil {
		r.logger.Error("Faced an error while tried to insert booking, err: ", err)
		return nil, errors.ErrFailedToInsert
	}
	return u, nil
}
