package models

import "time"

type Booking struct {
	BookingId   int       `json:"booking_id"`
	UserId      int       `json:"user_id"`
	DriverId    int       `json:"driver_id"`
	SeatsId     int       `json:"seats_id"`
	StatusId    int       `json:"status_id"`
	Price       float64   `json:"price"`
	StartCityId int       `json:"start_city_id"`
	EndCityId   int       `json:"end_city_id"`
	CreatedAt   time.Time `json:"created_at"`
	Active      bool      `json:"active"`
	UpdatedAt   time.Time `json:"updated_at"`
	CancelledAt time.Time `json:"cancelled_at"`
}
