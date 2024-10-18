package repository

import (
	"FinalProject/internal/models"
	"FinalProject/pkg/errors"
)

func (r *Repository) GetStatuses() ([]models.Status, error) {
	var statuses []models.Status
	query := `select * from statuses`
	err := r.db.Raw(query).Scan(&statuses).Error
	if err != nil {
		r.logger.Error("Faced an error while tried to select statuses, err: ", err)
		return nil, errors.ErrFailedToGet
	}
	return statuses, nil
}
