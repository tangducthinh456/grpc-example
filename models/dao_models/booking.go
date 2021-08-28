package daomodels

import (
	"gorm.io/gorm"
	"time"
)

type Booking struct {
	gorm.Model
	UUID       string    `json:"booking_code"`
	Date       time.Time `json:"booking_date"`
	CustomerId string    `json:"customer_id"`
	FlightId   string    `json:"flight_id"`
	Status     string    `json:"status"`
}
