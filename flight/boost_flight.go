package flight

import (
	"bookingFlight/flight/grpc"
	"bookingFlight/flight/services/repositories"
	"context"
	"bookingFlight/flight/api"
)

func BoostFlight() {
	repositories.InitDatabaseConnection()
	repositories.Flight().MigrateDataModel(context.Background())

	go func() {
		grpc.BoostGrpc()
	}()

	api.ApiBoost()

}
