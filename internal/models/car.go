package models

import (
	"time"
)

type Car struct {
	CarId     int       `json:"CarId"`
	Model     string    `json:"Model"`
	Mark      string    `json:"Mark"`
	Autobody  string    `json:"Autobody"`
	CarNumber string    `json:"car_number"`
	Seats     string    `json:"Seats"`
	UserId    int       `json:"user_id"`
	CreatedAt time.Time `json:"CreatedAt"`
	Active    bool      `json:"Active"`
	UpdatedAt time.Time `json:"UpdatedAt"`
	DeletedAt time.Time `json:"DeletedAt"`
}
