package convert

import (
	daomodels "bookingFlight/models/dao_models"
	dtomodels "bookingFlight/models/dto_models"
	"bookingFlight/pb"
)


type convert interface {
	PbToDto(*pb.Customer) dtomodels.Customer
	DtoToPb(dtomodels.Customer) *pb.Customer
	DtoToDao(dtomodels.Customer) daomodels.Customer
	DaoToDto(daomodels.Customer) dtomodels.Customer
}

type convertService struct {}

func (c *convertService) PbToDto(cusPb *pb.Customer) dtomodels.Customer {
	cusDto := dtomodels.Customer{
		ID: cusPb.GetId(),
		Name: cusPb.GetName(),
		Dob: cusPb.GetDob(),
		Email: cusPb.GetEmail(),
		Address: cusPb.GetAddress(),
		Phone: cusPb.GetPhone(),
	}
	return cusDto
}

func (c *convertService) DtoToPb(cusDto dtomodels.Customer) *pb.Customer {
	cusPb := &pb.Customer{
		Id: cusDto.ID,
		Name: cusDto.Name,
		Dob: cusDto.Dob,
		Address: cusDto.Address,
		Email: cusDto.Email,
		Phone: cusDto.Phone,
	}
	return cusPb
}

func (c *convertService) DtoToDao(cusDto dtomodels.Customer) daomodels.Customer {
	cusDao := daomodels.Customer{
		UUID: cusDto.ID,
		Name: cusDto.Name,
		Address: cusDto.Address,
		PhoneNumber: cusDto.Phone,
		Email: cusDto.Email,
		Dob: cusDto.Dob,
	}

	return cusDao
}

func (c *convertService) DaoToDto(cusDao daomodels.Customer) dtomodels.Customer {
	cusDto := dtomodels.Customer{
		ID: cusDao.UUID,
		Name: cusDao.Name,
		Dob: cusDao.Dob,
		Address: cusDao.Address,
		Email: cusDao.Email,
		Phone: cusDao.PhoneNumber,
	}
	return cusDto
}

var conv convert = &convertService{}

func Customer() convert {
	return conv
} 