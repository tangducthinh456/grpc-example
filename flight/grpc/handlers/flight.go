package handlers

import (
	"bookingFlight/flight/services/convert"
	"bookingFlight/flight/services/gendata"
	"bookingFlight/flight/services/repositories"
	"bookingFlight/log"
	"bookingFlight/pb"
	"context"
	"fmt"
)

type flight interface {
	CreateFlight(ctx context.Context, in *pb.Flight) (*pb.Flight, error)
	UpdateFlight(ctx context.Context, in *pb.RequestUpdateFlight) (*pb.Flight, error)
	SearchFlight(ctx context.Context, in *pb.RequestSearchFlight) (*pb.Flight, error)
}

type FlightHandler struct {
	pb.UnimplementedFlightServiceServer
}

// var flightH flight = &flightHandler{}

// func Flight() flight {
// 	return flightH
// }

func (c *FlightHandler) CreateFlight(ctx context.Context, in *pb.Flight) (*pb.Flight, error) {
	log.Info("gRPC call create flight")
	flightDto := convert.Flight().PbToDto(in)
	flightDao, err := convert.Flight().DtoToDao(flightDto)
	if err != nil {
		return nil, err
	}

	flightDao.UUID = gendata.GenUUID()
	flightDao.AvailableSlot = 100

	err = repositories.Flight().CreateFlight(ctx, flightDao)
	if err != nil {
		log.Error(fmt.Sprintf("gRPC call create flight fail : %s", err.Error()))
		return nil, err
	}
	flightDto = convert.Flight().DaoToDto(*flightDao)
	fliPb := convert.Flight().DtoToPb(flightDto)
	return fliPb, nil
}

func (c *FlightHandler) UpdateFlight(ctx context.Context, in *pb.RequestUpdateFlight) (*pb.Flight, error) {
	log.Info("gRPC call update flight")
	flightDto := convert.Flight().PbToDto(in.Flight)
	flightDao, err := convert.Flight().DtoToDao(flightDto)
	if err != nil {
		return nil, err
	}

	err = repositories.Flight().UpdateFlight(ctx, in.Id, flightDao)
	if err != nil {
		log.Error(fmt.Sprintf("gRPC call update flight fail : %s", err.Error()))
		return nil, err
	}
	flightDto = convert.Flight().DaoToDto(*flightDao)
	fliPb := convert.Flight().DtoToPb(flightDto)
	return fliPb, nil
}

func (c *FlightHandler) SearchFlight(ctx context.Context, in *pb.RequestSearchFlight) (*pb.Flight, error) {
	log.Info("gRPC call search fli")

	fliDao, err := repositories.Flight().GetFlightByID(ctx, in.Id)
	if err != nil {
		return nil, err
	}
	fliDto := convert.Flight().DaoToDto(*fliDao)
	fliPb := convert.Flight().DtoToPb(fliDto)

	return fliPb, nil
}
