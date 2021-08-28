package handler

import (
	"bookingFlight/flight/api/models"
	"bookingFlight/flight/services/convert"
	"bookingFlight/log"
	dtomodels "bookingFlight/models/dto_models"
	"bookingFlight/pb"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type FlightHandler interface {
	HandleCreateFlight(c *gin.Context)
	HandleUpdateFlight(c *gin.Context)
	HandleSearchFlight(c *gin.Context)
	// BookingHistory(c *gin.Context)
}

type flightHandler struct {
	flightClient pb.FlightServiceClient
}

var cus FlightHandler

func Flight() FlightHandler {
	return cus
}

func InitFlightHandler(flightClient *pb.FlightServiceClient) {
	cus = &flightHandler{
		flightClient: *flightClient,
	}
}

func (cus *flightHandler) HandleCreateFlight(c *gin.Context) {

	log.Info("handle request create flight")

	req := &dtomodels.Flight{}

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

	reqPb := convert.Flight().DtoToPb(*req)
	cusPb, err := cus.flightClient.CreateFlight(c, reqPb)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    []string{err.Error()},
		})
		return
	}

	cusDto := convert.Flight().PbToDto(cusPb)

	c.JSON(http.StatusCreated, models.Response{
		StatusCode: http.StatusCreated,
		Payload:    cusDto,
	})
	return

}

func (cus *flightHandler) HandleUpdateFlight(c *gin.Context) {
	log.Info("handle request update flight")

	req := &dtomodels.Flight{}

	flightId := c.Param("FlightId")

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

	reqPb := convert.Flight().DtoToPb(*req)
	reqUpdate := &pb.RequestUpdateFlight{
		Id:     flightId,
		Flight: reqPb,
	}
	cusPb, err := cus.flightClient.UpdateFlight(c, reqUpdate)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    []string{err.Error()},
		})
		return
	}

	cusDto := convert.Flight().PbToDto(cusPb)

	c.JSON(http.StatusCreated, models.Response{
		StatusCode: http.StatusCreated,
		Payload:    cusDto,
	})
	return
}

func (cus *flightHandler) HandleSearchFlight(c *gin.Context) {
	log.Info("handle request update password")

	flightId := c.Param("FlightId")

	reqChange := &pb.RequestSearchFlight{
		Id: flightId,
	}
	fliPb, err := cus.flightClient.SearchFlight(c, reqChange)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message:    []string{err.Error()},
		})
		c.Abort()
		return
	}

	fliDto := convert.Flight().PbToDto(fliPb)

	c.JSON(http.StatusOK, models.Response{
		StatusCode: http.StatusOK,
		Payload:    fliDto,
	})
}
