package models

type Seat struct {
	SeatId      int  `json:"seat_id"`
	CarId       int  `json:"car_id"`
	SeatNumber  int  `json:"seat_number"`
	IsAvailible bool `json:"is_availible"`
}
type Seats struct {
	CarId int   `json:"car_id"`
	Seats []int `json:"seats"`
}
