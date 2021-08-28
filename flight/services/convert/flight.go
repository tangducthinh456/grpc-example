package convert

import (
	daomodels "bookingFlight/models/dao_models"
	dtomodels "bookingFlight/models/dto_models"
	"bookingFlight/pb"
	"fmt"
	"time"
)

type convert interface {
	PbToDto(*pb.Flight) dtomodels.Flight
	DtoToPb(dtomodels.Flight) *pb.Flight
	DtoToDao(dtomodels.Flight) (*daomodels.Flight, error)
	DaoToDto(daomodels.Flight) dtomodels.Flight
}

type convertService struct{}

func (c *convertService) PbToDto(fliPb *pb.Flight) dtomodels.Flight {
	fliDto := dtomodels.Flight{
		ID:          fliPb.GetId(),
		From:        fliPb.GetFrom(),
		To:          fliPb.GetTo(),
		FlightPlane: fliPb.GetFlightPlane(),
		DepartDate:  fliPb.GetDepartDate(),
		DepartTime:  fliPb.GetDepartTime(),
		Status:      fliPb.Status,
	}
	return fliDto
}

func (c *convertService) DtoToPb(fliDto dtomodels.Flight) *pb.Flight {
	fliPb := &pb.Flight{
		Id:          fliDto.ID,
		From:        fliDto.From,
		To:          fliDto.To,
		FlightPlane: fliDto.FlightPlane,
		DepartDate:  fliDto.DepartDate,
		DepartTime:  fliDto.DepartTime,
		Status:      fliDto.Status,
	}
	return fliPb
}

func (c *convertService) DtoToDao(fliDto dtomodels.Flight) (*daomodels.Flight, error) {
	timeFormat := fmt.Sprintf("%sT%sZ", fliDto.DepartDate, fliDto.DepartTime)
	timeFly, err := time.Parse(time.RFC3339, timeFormat)
	if err != nil {
		return nil, err
	}
	fliDao := &daomodels.Flight{
		UUID:   fliDto.ID,
		Name:   fliDto.FlightPlane,
		From:   fliDto.From,
		To:     fliDto.To,
		Status: fliDto.Status,
		Date:   timeFly,
	}

	return fliDao, nil
}

func (c *convertService) DaoToDto(fliDao daomodels.Flight) dtomodels.Flight {
	fliDto := dtomodels.Flight{
		ID:          fliDao.UUID,
		FlightPlane: fliDao.Name,
		From:        fliDao.From,
		To:          fliDao.To,
		Status:      fliDao.Status,
		DepartDate:  fliDao.Date.Format(time.RFC3339)[:10],
		DepartTime:  fliDao.Date.Format(time.RFC3339)[11:19],
	}
	return fliDto
}

var conv convert = &convertService{}

func Flight() convert {
	return conv
}
