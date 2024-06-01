package entity

import "time"

type Ticket struct {
	Company       string
	Departure     time.Time
	Arrival       time.Time
	Origin        string // TODO: make an enum of cities
	Destination   string
	Price         int64
	LeftSeats     int32
	VehicleDetail string
}
