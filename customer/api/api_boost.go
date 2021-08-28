package api

import (
	"bookingFlight/config"
	"bookingFlight/customer/api/handler"
	"bookingFlight/customer/api/router"
	"bookingFlight/log"
	"bookingFlight/pb"
	"context"
	"errors"
	"google.golang.org/grpc"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func ApiBoost() {

	log.Info("Initialize api service")

	cusomerConn, err := grpc.Dial(":"+config.CustomerConfig().GRPC.Port, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	customerClient := pb.NewCustomerServiceClient(cusomerConn)
	handler.InitCustomerHandler(&customerClient)

	router.Init()

	var httpServer = http.Server{
		Addr:    ":" + config.CustomerConfig().API.Port,
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
