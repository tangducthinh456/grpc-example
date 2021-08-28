package booking

import (
	"bookingFlight/booking/grpc"
	"bookingFlight/booking/services/repositories"
	"context"
	"bookingFlight/flight/api"
)

func BoostBooking() {

	repositories.InitDatabaseConnection()
	repositories.Booking().MigrateDataModel(context.Background())

	go func() {
		grpc.BoostGrpc()
	}()

	api.ApiBoost()

}
