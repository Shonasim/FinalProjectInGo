package repository

import (
	"FinalProject/internal/models"
	"fmt"
)

func (r *Repository) AddUser(u *models.User) (*models.User, error) {
	result := r.db.Create(&u)
	if result.Error != nil {
		return nil, fmt.Errorf("Failed to add user: %v\n", result.Error)
	}
	return u, nil
}
