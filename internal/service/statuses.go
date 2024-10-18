package service

import "FinalProject/internal/models"

func (s *Service) GetStatuses() ([]models.Status, error) {
	return s.Repository.GetStatuses()
}
