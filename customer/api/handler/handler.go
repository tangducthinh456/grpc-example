package handler

import (
	"bookingFlight/customer/api/models"
	"bookingFlight/log"
	dtomodels "bookingFlight/models/dto_models"
	"bookingFlight/pb"
	"bookingFlight/customer/services/convert"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CustomerHandler interface {
	HandleCreateCustomer(c *gin.Context)
	HandleUpdateCustomer(c *gin.Context)
	HandleChangePassword(c *gin.Context)
	HandleSearchCustomer(c *gin.Context)

	// BookingHistory(c *gin.Context)
}

type customerHandler struct{
	customerClient pb.CustomerServiceClient
}

var cus CustomerHandler

func Customer() CustomerHandler {
	 return cus
}

func InitCustomerHandler(customerClient *pb.CustomerServiceClient) {
	cus = &customerHandler{
		customerClient: *customerClient,
	}
}

func (cus *customerHandler) HandleCreateCustomer(c *gin.Context) {

	log.Info("handle request create customer")
	
	req := &dtomodels.Customer{}
	
	if err := c.ShouldBindJSON(&req); err != nil {
		if validateErrors, ok := err.(validator.ValidationErrors); ok {
			errMessages := make([]string, 0)
			for _, v := range validateErrors {
				errMessages = append(errMessages, v.Error())
				log.Error(fmt.Sprintf("Binding request error : %s",v.Error()))
			}

			c.AbortWithStatusJSON(http.StatusBadRequest, models.Response{
				StatusCode: http.StatusBadRequest,
				Message: errMessages,
			})
			return
		}
		log.Error(fmt.Sprintf("Binding request error : %s",err.Error()))

		c.AbortWithStatusJSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message: []string{err.Error()},
		})
		return
	}

	reqPb := convert.Customer().DtoToPb(*req)
	cusPb, err := cus.customerClient.CreateCustomer(c, reqPb)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message: []string{err.Error()},
		})
		return
	}

	cusDto := convert.Customer().PbToDto(cusPb)

	c.JSON(http.StatusCreated, models.Response{
		StatusCode: http.StatusCreated,
		Payload: cusDto,
	})
	return

}

func (cus *customerHandler) HandleUpdateCustomer(c *gin.Context) {
	log.Info("handle request update customer")
	
	req := &dtomodels.Customer{}

	customerId := c.Param("CustomerId")
	
	if err := c.ShouldBindJSON(&req); err != nil {
		if validateErrors, ok := err.(validator.ValidationErrors); ok {
			errMessages := make([]string, 0)
			for _, v := range validateErrors {
				errMessages = append(errMessages, v.Error())
				log.Error(fmt.Sprintf("Binding request error : %s",v.Error()))
			}

			c.AbortWithStatusJSON(http.StatusBadRequest, models.Response{
				StatusCode: http.StatusBadRequest,
				Message: errMessages,
			})
			return
		}
		log.Error(fmt.Sprintf("Binding request error : %s",err.Error()))

		c.AbortWithStatusJSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message: []string{err.Error()},
		})
		return
	}

	reqPb := convert.Customer().DtoToPb(*req)
	reqUpdate := &pb.RequestUpdateCustomer{
		Id: customerId,
		Customer: reqPb,
	}
	cusPb, err := cus.customerClient.UpdateCustomer(c, reqUpdate)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message: []string{err.Error()},
		})
		return
	}

	cusDto := convert.Customer().PbToDto(cusPb)

	c.JSON(http.StatusCreated, models.Response{
		StatusCode: http.StatusCreated,
		Payload: cusDto,
	})
	return
}

func (cus *customerHandler) HandleChangePassword(c *gin.Context) {
	log.Info("handle request update password")
	
	req := &dtomodels.RequestChangePassword{}
	customerId := c.Param("CustomerId")
	if err := c.ShouldBindJSON(&req); err != nil {
		if validateErrors, ok := err.(validator.ValidationErrors); ok {
			errMessages := make([]string, 0)
			for _, v := range validateErrors {
				errMessages = append(errMessages, v.Error())
				log.Error(fmt.Sprintf("Binding request error : %s",v.Error()))
			}

			c.AbortWithStatusJSON(http.StatusBadRequest, models.Response{
				StatusCode: http.StatusBadRequest,
				Message: errMessages,
			})
			return
		}
		log.Error(fmt.Sprintf("Binding request error : %s",err.Error()))

		c.AbortWithStatusJSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message: []string{err.Error()},
		})
		return
	}

	reqChange := &pb.RequestChangePasswordCustomer{
		Id: customerId,
		OldPassword: req.OldPassword,
		NewPassword: req.NewPassword,
	}
	_, err := cus.customerClient.ChangePassword(c, reqChange)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message: []string{err.Error()},
		})
		c.Abort()
		return 
	}

	c.JSON(http.StatusOK, models.Response{
		StatusCode: http.StatusOK,
		Payload: nil,
	})
}

func (cus *customerHandler) HandleSearchCustomer(c *gin.Context) {
	customerId := c.Param("CustomerId")
	reqSearch := pb.RequestSearchCustomer{
		Id: customerId,
	}
	customerPb, err := cus.customerClient.SearchCustomer(c, &reqSearch)
	if err != nil {
		log.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Response{
			StatusCode: http.StatusBadRequest,
			Message: []string{err.Error()},
		})
		c.Abort()
		return 
	}
	customerDto := convert.Customer().PbToDto(customerPb)
	c.JSON(http.StatusOK, models.Response{
		StatusCode: http.StatusOK,
		Payload: customerDto,
	})
}