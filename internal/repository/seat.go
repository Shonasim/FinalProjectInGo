package repository

import (
	"FinalProject/internal/models"
	"FinalProject/pkg/errors"
)

func (r *Repository) AddSeat(s *models.Seats) (*models.Seats, error) {
	for _, seat := range s.Seats {
		query := `insert into seats (car_id, seat_number) values(?,?)`
		err := r.db.Exec(query, s.CarId, seat).Error
		if err != nil {
			r.logger.Error("Faced an error while tried to insert seat, err: ", err)
			return nil, errors.ErrFailedToCreateSeat
		}
	}
	return s, nil
}

func (r *Repository) GetSeats(carId int) ([]models.Seat, error) {
	var seats []models.Seat
	query := `select * from seats where car_id = ? and available = true`
	err := r.db.Raw(query, carId).Scan(&seats).Error
	if err != nil {
		r.logger.Error("Faced an error while tried to get seats, err: ", err)
		return nil, errors.ErrFailedToGet
	}
	return seats, nil
}
