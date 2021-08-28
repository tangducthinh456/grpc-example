package handlers

import (
	"context"
	"bookingFlight/log"
	"bookingFlight/pb"
	"bookingFlight/customer/services/convert"
	"bookingFlight/customer/services/gendata"
	"bookingFlight/customer/services/repositories"
	"errors"
	"fmt"
)

type customer interface {
	CreateCustomer(ctx context.Context, in *pb.Customer) (*pb.Customer, error)
	UpdateCustomer(ctx context.Context, in *pb.RequestUpdateCustomer) (*pb.Customer, error)
	ChangePasswordCustomer(ctx context.Context, in *pb.RequestChangePasswordCustomer) (*pb.Empty, error)
	SearchCustomer(ctx context.Context, in *pb.RequestSearchCustomer) (*pb.Customer, error)
	// Booking history
}

type CustomerHandler struct {
	pb.UnimplementedCustomerServiceServer
}

// var customerH customer = &customerHandler{}

// func Customer() customer {
// 	return customerH
// }

func (c *CustomerHandler) CreateCustomer(ctx context.Context, in *pb.Customer) (*pb.Customer, error) {
	log.Info("gRPC call create customer")
	customerDto := convert.Customer().PbToDto(in)
	customerDao := convert.Customer().DtoToDao(customerDto)
	customerDao.UUID = gendata.GenUUID()

	err := repositories.Customer().CreateCustomer(ctx, &customerDao)
	if err != nil {
		log.Error(fmt.Sprintf("gRPC call create customer fail : %s", err.Error()))
		return nil, err
	}
	customerDto = convert.Customer().DaoToDto(customerDao)
	cusPb := convert.Customer().DtoToPb(customerDto)
	return cusPb, nil
}

func (c *CustomerHandler) UpdateCustomer(ctx context.Context, in *pb.RequestUpdateCustomer) (*pb.Customer, error) {
	log.Info("gRPC call update customer")
	customerDto := convert.Customer().PbToDto(in.Customer)
	customerDao := convert.Customer().DtoToDao(customerDto)

	err := repositories.Customer().UpdateCustomer(ctx, in.Id, &customerDao)
	if err != nil {
		log.Error(fmt.Sprintf("gRPC call update customer fail : %s", err.Error()))
		return nil, err
	}
	customerDto = convert.Customer().DaoToDto(customerDao)
	cusPb := convert.Customer().DtoToPb(customerDto)
	return cusPb, nil
}

func (c *CustomerHandler) ChangePassword(ctx context.Context, in *pb.RequestChangePasswordCustomer) (*pb.Empty, error) {
	log.Info("gRPC call update password")

	encryptPass, err := gendata.HashPassword(in.NewPassword)
	if err != nil {
		log.Error(fmt.Sprintf("generate hash password fail : %s", err.Error()))
		return nil, err
	}

	customerDao, err := repositories.Customer().GetCustomerByID(ctx, in.Id)
	if err != nil {
		log.Error(fmt.Sprintf("get customer fail : %s", err.Error()))
		return nil, err
	}

	if customerDao.Password == "" {

		customerDao.Password = encryptPass
		err = repositories.Customer().UpdateCustomer(ctx, in.Id, customerDao)
		if err != nil {
			log.Error(fmt.Sprintf("update customer fail : %s", err.Error()))
			return nil, err
		}
		return &pb.Empty{}, nil
	}

	if in.OldPassword == "" {
		log.Error(fmt.Sprintf("old password is empty : %s", err.Error()))
		err = errors.New("old password is empty")
		return nil, err
	}

	if !gendata.CheckPasswordHash(in.OldPassword, customerDao.Password) {

		log.Error("old password does not match")
		err = errors.New("old password does not match")
		return nil, err

	}

	customerDao.Password = encryptPass
	err = repositories.Customer().UpdateCustomer(ctx, in.Id, customerDao)
	if err != nil {
		log.Error(fmt.Sprintf("update customer fail : %s", err.Error()))
		return nil, err
	}
	return &pb.Empty{}, nil
}

func (c *CustomerHandler) SearchCustomer(ctx context.Context, in *pb.RequestSearchCustomer) (*pb.Customer, error) {
	customerDao, err := repositories.Customer().GetCustomerByID(ctx, in.Id)
	if err != nil {
		return nil, err
	}
	customerDto := convert.Customer().DaoToDto(*customerDao)
	customerPb := convert.Customer().DtoToPb(customerDto)

	return customerPb, nil
}
