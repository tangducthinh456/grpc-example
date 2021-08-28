package api

import (
	"bookingFlight/config"
	"bookingFlight/flight/api/handler"
	"bookingFlight/flight/api/router"
	"bookingFlight/pb"
	"context"

	// "flight/config"
	"bookingFlight/log"
	"errors"
	// "fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	// "golang.org/x/net/http2"
	"google.golang.org/grpc"
)

func ApiBoost() {

	log.Info("Initialize api service")

	cusomerConn, err := grpc.Dial(":"+config.FlightConfig().GRPC.Port, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	flightClient := pb.NewFlightServiceClient(cusomerConn)
	handler.InitFlightHandler(&flightClient)

	router.Init()

	var httpServer = http.Server{
		Addr:    ":" + config.FlightConfig().API.Port,
		Handler: router.Router(),
	}

	// var http2Server = http2.Server{}
	// err = http2.ConfigureServer(&httpServer, &http2Server)
	// if err != nil {
	// 	log.Error(fmt.Sprintf("Config server http2 fail : %s", err.Error()))
	// }

	go func() {
		if err := httpServer.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Info("Start serving")
		}
	}()

	// graceful shutdown
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Info("Server forced to shutdown:")
		log.Info(err.Error())
	}
	log.Info("Server exiting")
}
