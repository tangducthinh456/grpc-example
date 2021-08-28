package main

import (
	"bookingFlight/customer"
	"bookingFlight/config"
	"bookingFlight/flight"
	"bookingFlight/booking"
)

func main(){
	config.Init()

	go func(){
		customer.BoostCustomer()
	}()

	go func(){
		flight.BoostFlight()	
	}()

	booking.BoostBooking()
	
}