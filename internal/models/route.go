package models

import "time"

type Route struct {
	RouteId   int       `json:"route_id"`
	FromCity  int       `json:"from_city"`
	ToCity    int       `json:"to_city"`
	Price     float64   `json:"price"`
	Date      time.Time `json:"date"`
	DriverId  int       `json:"driver_id"`
	CarId     int       `json:"car_id"`
	CreatedAt time.Time `json:"created_at"`
	Active    bool      `json:"active"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
