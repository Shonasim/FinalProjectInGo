package repository

import (
	"FinalProject/internal/models"
	"FinalProject/pkg/errors"
	"time"
)

func (r *Repository) GetCar(userId int) ([]models.Car, error) {
	query := `SELECT * FROM cars WHERE user_id = ?`

	var car []models.Car
	err := r.db.Raw(query, userId).Scan(&car).Error
	if err != nil {
		r.logger.Error("Faced an error while trying to select car from table, err: ", err)
		return nil, err // Возвращаем nil для автомобиля и ошибку
	}
	return car, nil // Возвращаем указатель на автомобиль
}
func (r *Repository) AddCar(u *models.Car) (*models.Car, error) {
	query := `INSERT INTO cars
    (model,mark,autobody,car_number,seats,user_id, created_at)
	VALUES (?,?,?,?,?,?,?)`
	err := r.db.Exec(query, u.Model, u.Mark, u.Autobody, u.CarNumber, u.Seats, u.UserId, time.Now()).Error
	if err != nil {
		r.logger.Error("Faced an error while tried to insert car, err: ", err)
		return nil, errors.ErrFailedCreateCar
	}
	return u, nil
}
