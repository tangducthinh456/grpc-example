package customer

import (
	"context"
	"bookingFlight/customer/grpc"
	"bookingFlight/customer/services/repositories"
	"bookingFlight/customer/api"
)

func BoostCustomer() {
	 
	repositories.InitDatabaseConnection()
	repositories.Customer().MigrateDataModel(context.Background())
	
	go func(){
		grpc.BoostGrpc()
	}()
	
	api.ApiBoost()

}
