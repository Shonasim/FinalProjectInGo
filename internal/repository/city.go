package repository

import (
	"FinalProject/internal/models"
	errors2 "FinalProject/pkg/errors"
)

func (r *Repository) GetCitiesList() ([]models.City, error) {
	var cities []models.City
	query := `select * from cities`
	err := r.db.Raw(query).Scan(&cities).Error
	if err != nil {
		r.logger.Error("Faced an error while tried to select cities, err: ", err)
		return nil, errors2.ErrGetCities
	}
	return cities, nil
}
