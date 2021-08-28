package handler

import (
	"bookingFlight/booking/api/models"
	"bookingFlight/booking/services/convert"
	"bookingFlight/booking/services/repositories"
	"bookingFlight/log"

	// dtomodels "bookingFlight/models/dto_models"
	daomodels "bookingFlight/models/dao_models"
	"bookingFlight/pb"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type BookingHandler interface {
	HandleCreateBooking(c *gin.Context)
	HandleCancelBooking(c *gin.Context)
	HandleViewBooking(c *gin.Context)
	// BookingHistory(c *gin.Context)
}

type bookingHandler struct {
	bookingClient pb.BookingServiceClient
}

var cus BookingHandler

func Booking() BookingHandler {
	return cus
}

func InitBookingHandler(bookingClient *pb.BookingServiceClient) {
	cus = &bookingHandler{
		bookingClient: *bookingClient,
	}
}

func (cus *bookingHandler) HandleCreateBooking(c *gin.Context) {

	log.Info("handle request create booking")

	req := &daomodels.Booking{}

	if err := c.ShouldBindJSON(&req); err != nil {
		if validateErrors, ok := err.(validator.ValidationErrors); ok {
			errMessages := make([]string, 0)
			for _, v := range validateErrors {
				errMessages = append(errMessages, v.Error())
				log.Error(fmt.Sprintf("Binding request error : %s", v.Error()))
			}

			c.AbortWithStatusJSON(http.StatusBadRequest, models.Response{
				StatusCode: http.StatusBadRequest,
				Message:    errMessages,
			})
			return
		}
		log.Error(fmt.Sprintf("Binding request error : %s", err.Error()))

		c.AbortWithStatusJSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    []string{err.Error()},
		})
		return
	}

	reqPb := convert.Booking().DaoToPb(*req)
	cusPb, err := cus.bookingClient.CreateBooking(c, &pb.RequestBooking{
		FlightId:   reqPb.FlightId,
		CustomerId: reqPb.CustomerId,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    []string{err.Error()},
		})
		return
	}

	cusDao, err := convert.Booking().PbToDao(cusPb)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    []string{err.Error()},
		})
		return
	}

	c.JSON(http.StatusCreated, models.Response{
		StatusCode: http.StatusCreated,
		Payload:    cusDao,
	})
	return

}

func (cus *bookingHandler) HandleViewBooking(c *gin.Context) {
	log.Info("handle request update password")

	bookingId := c.Param("BookingId")

	reqChange := &pb.RequestViewBooking{
		Id: bookingId,
	}
	fliPb, err := cus.bookingClient.ViewBooking(c, reqChange)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    []string{err.Error()},
		})
		c.Abort()
		return
	}

	fliDto, _ := convert.Booking().PbToDao(fliPb)

	c.JSON(http.StatusOK, models.Response{
		StatusCode: http.StatusOK,
		Payload:    fliDto,
	})
}

func (cus *bookingHandler) HandleCancelBooking(c *gin.Context) {
	id := c.Param("BookingId")
	err := repositories.Booking().DeleteBooking(c, id)
	if err != nil {
		log.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    []string{err.Error()},
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, models.Response{
		StatusCode: http.StatusOK,
		Payload:    nil,
	})
}
