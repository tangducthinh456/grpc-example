package handlers

import (
	// daomodels "booking/models/dao_models"
	"bookingFlight/booking/services/convert"
	"bookingFlight/booking/services/gendata"
	"bookingFlight/booking/services/repositories"
	"bookingFlight/config"
	"bookingFlight/log"
	"bookingFlight/pb"
	"context"
	"errors"
	"fmt"
	"time"

	"google.golang.org/grpc"
)

type booking interface {
	CreateBooking(ctx context.Context, in *pb.Booking) (*pb.Booking, error)
	ViewBooking(ctx context.Context, in *pb.RequestViewBooking) (*pb.Booking, error)
	GetBookingByUser(ctx context.Context, id string) ([]pb.Booking, error)
	CancelBooking(ctx context.Context, in *pb.RequestCancelBooking) (*pb.Booking, error)
}

type BookingHandler struct {
	pb.UnimplementedBookingServiceServer
}

// var bookingH booking = &bookingHandler{}

// func Booking() booking {
// 	return bookingH
// }

func (c *BookingHandler) CreateBooking(ctx context.Context, in *pb.Booking) (*pb.Booking, error) {
	log.Info("gRPC call create booking")

	reqSearchFli := &pb.RequestSearchFlight{
		Id: in.FlightId,
	}

	flightConn, err := grpc.Dial(":"+config.FlightConfig().GRPC.Port, grpc.WithInsecure())
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	// if fliPb.
	a := pb.NewFlightServiceClient(flightConn)
	flightPb, err := a.SearchFlight(ctx, reqSearchFli)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	if flightPb.Slot == 0 {
		err = errors.New("plane is full, can not book any more")
		log.Error(err.Error())
		return nil, err
	}

	flightPb.Slot = flightPb.Slot - 1

	reqUpdate := pb.RequestUpdateFlight{
		Id:     flightPb.Id,
		Flight: flightPb,
	}
	_, err = a.UpdateFlight(ctx, &reqUpdate)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	bookingDao, err := convert.Booking().PbToDao(in)
	if err != nil {
		log.Error("gRPC create booking failed : " + err.Error())
		return nil, err
	}
	bookingDao.UUID = gendata.GenUUID()

	err = repositories.Booking().CreateBooking(ctx, bookingDao)
	if err != nil {
		log.Error(fmt.Sprintf("gRPC call create booking fail : %s", err.Error()))
		return nil, err
	}

	boPb := convert.Booking().DaoToPb(*bookingDao)

	return boPb, nil
}

func (c *BookingHandler) ViewBooking(ctx context.Context, in *pb.RequestViewBooking) (*pb.BookingAssociate, error) {
	log.Info("gRPC call view book")

	bookingDao, err := repositories.Booking().GetBookingByID(ctx, in.Id)
	if err != nil {
		return nil, err
	}

	flightConn, err := grpc.Dial(":"+config.FlightConfig().GRPC.Port, grpc.WithInsecure())
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	cusConn, err := grpc.Dial(":"+config.FlightConfig().GRPC.Port, grpc.WithInsecure())
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	a := pb.NewCustomerServiceClient(cusConn)

	cusPb, err := a.SearchCustomer(ctx, &pb.RequestSearchCustomer{
		Id: bookingDao.CustomerId,
	})

	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	b := pb.NewFlightServiceClient(flightConn)
	fliPb, err := b.SearchFlight(ctx, &pb.RequestSearchFlight{
		Id: bookingDao.FlightId,
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	res := &pb.BookingAssociate{
		BookingCode: bookingDao.UUID,
		BookingDate: bookingDao.Date.Format(time.RFC3339)[:10],
		Customer:    cusPb,
		Flight:      fliPb,
	}

	return res, nil
}


func (c *BookingHandler) CancelBooking(ctx context.Context,req  pb.RequestCancelBooking) (*pb.Empty, error){
	err := repositories.Booking().DeleteBooking(ctx, req.Id)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	return &pb.Empty{}, nil
}

func (c *BookingHandler) GetBookingByUser(ctx context.Context, req pb.RequestViewBooking) (*pb.ListBooking, error){
	list, err := repositories.Booking().GetListBookingByCustomer(ctx, req.Id)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	_ = list
	//not implement yet
	return nil, nil
}