package service

import "FinalProject/internal/models"

func (s *Service) GetServiceList() ([]models.City, error) {
	return s.Repository.GetCitiesList()
}
