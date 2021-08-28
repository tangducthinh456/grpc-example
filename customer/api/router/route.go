package router

import (
	"bookingFlight/customer/api/handler"
	"bookingFlight/log"

	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func Init() {
	log.Info("Initialize router")

	router = gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.POST("/", handler.Customer().HandleCreateCustomer) 
	router.PUT("/:CustomerId", handler.Customer().HandleUpdateCustomer)
	router.PUT("/:CustomerId/password", handler.Customer().HandleChangePassword)
	router.GET("/:CustomerId", handler.Customer().HandleSearchCustomer)
	// router.GET("/booking", handler.Customer().HandleChangePassword)
}

func Router() *gin.Engine {
	return router
}
