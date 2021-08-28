package grpc

import (
	"bookingFlight/config"
	"bookingFlight/customer/grpc/handlers"
	"bookingFlight/log"
	"bookingFlight/pb"
	"fmt"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
	// "google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
)

func BoostGrpc() {

	log.Info("Start grpc services")

	listen, err := net.Listen("tcp", ":" + config.CustomerConfig().GRPC.Port)
	if err != nil {
		log.Error(fmt.Sprintf("net listen failed : %s", err.Error()))
		panic(err)
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_recovery.UnaryServerInterceptor(),
		),
		))

	reflection.Register(s)

	pb.RegisterCustomerServiceServer(s, &handlers.CustomerHandler{})

	s.Serve(listen)
}
