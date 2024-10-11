package repository

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Repository struct {
	db     *gorm.DB
	logger *logrus.Logger
}

func NewRepository(db *gorm.DB, log *logrus.Logger) *Repository {
	return &Repository{db: db, logger: log}
}
