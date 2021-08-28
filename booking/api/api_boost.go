package api

import (
	"bookingFlight/booking/api/handler"
	"bookingFlight/booking/api/router"
	"bookingFlight/config"
	"bookingFlight/pb"
	"context"

	// "booking/config"
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

	cusomerConn, err := grpc.Dial(":"+config.BookingConfig().GRPC.Port, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	bookingClient := pb.NewBookingServiceClient(cusomerConn)
	handler.InitBookingHandler(&bookingClient)

	router.Init()

	var httpServer = http.Server{
		Addr:    ":" + config.BookingConfig().API.Host,
		Handler: router.Router(),
	}

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
