package service

import "FinalProject/internal/repository"

type Service struct {
	Repository repository.Repository
}

func NewService(r repository.Repository) *Service {
	return &Service{Repository: r}
}
