package convert

import (
	daomodels "bookingFlight/models/dao_models"
	// dtomodels "bookingFlight/models/dto_models"

	"bookingFlight/pb"
	"fmt"
	"time"
)

type convert interface {
	PbToDao(*pb.Booking) (*daomodels.Booking, error)
	DaoToPb(daomodels.Booking) *pb.Booking
	// DtoToDao(dtomodels.Booking) (*daomodels.Booking, error)
	// DaoToDto(daomodels.Booking) dtomodels.Booking
}

type convertService struct{}

func (c *convertService) PbToDao(boPb *pb.Booking) (*daomodels.Booking, error) {
	timeFormat := fmt.Sprintf("%sT00:00:00Z", boPb.BookingDate)
	t, err := time.Parse(time.RFC3339, timeFormat)
	if err != nil {
		return nil, err
	}

	bookingDao := daomodels.Booking{
		UUID:       boPb.BookingCode,
		Date:       t,
		CustomerId: boPb.CustomerId,
		FlightId:   boPb.FlightId,
		Status:     boPb.Status,
	}
	return &bookingDao, nil
}

func (c *convertService) DaoToPb(bookingDao daomodels.Booking) *pb.Booking {
	boPb := &pb.Booking{
		BookingCode: bookingDao.UUID,
		BookingDate: bookingDao.Date.Format(time.RFC3339)[:10],
		CustomerId:  bookingDao.CustomerId,
		FlightId:    bookingDao.FlightId,
		Status:      bookingDao.Status,
	}
	return boPb
}

// func (c *convertService) DtoToDao(fliDto dtomodels.Booking) (*daomodels.Booking, error) {
// 	timeFormat := fmt.Sprintf("%sT%sZ", fliDto.DepartDate, fliDto.DepartTime)
// 	timeFly, err := time.Parse(time.RFC3339, timeFormat)
// 	if err != nil {
// 		return nil, err
// 	}
// 	fliDao := &daomodels.Booking{
// 		UUID:   fliDto.ID,
// 		Name:   fliDto.BookingPlane,
// 		From:   fliDto.From,
// 		To:     fliDto.To,
// 		Status: fliDto.Status,
// 		Date:   timeFly,
// 	}

// 	return fliDao, nil
// }

// func (c *convertService) DaoToDto(fliDao daomodels.Booking) dtomodels.Booking {
// 	fliDto := dtomodels.Booking{
// 		ID:          fliDao.UUID,
// 		BookingPlane: fliDao.Name,
// 		From:        fliDao.From,
// 		To:          fliDao.To,
// 		Status:      fliDao.Status,
// 		DepartDate:  fliDao.Date.Format("2016-01-30"),
// 		DepartTime:  fliDao.Date.Format("07:00:00"),
// 	}
// 	return fliDto
// }

var conv convert = &convertService{}

func Booking() convert {
	return conv
}
