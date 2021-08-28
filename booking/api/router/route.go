package router

import (
	"bookingFlight/booking/api/handler"
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

	router.POST("/", handler.Booking().HandleCreateBooking) 
	router.PUT("/:BookingId", handler.Booking().HandleCancelBooking)
	router.GET("/:BookingId", handler.Booking().HandleViewBooking)
	// router.GET("/booking", handler.Flight().HandleChangePassword)
}

func Router() *gin.Engine {
	return router
}
