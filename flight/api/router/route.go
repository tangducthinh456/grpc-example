package router

import (
	"bookingFlight/flight/api/handler"
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

	router.POST("/", handler.Flight().HandleCreateFlight) 
	router.PUT("/:FlightId", handler.Flight().HandleUpdateFlight)
	router.GET("/:FlightId", handler.Flight().HandleSearchFlight)
	// router.GET("/booking", handler.Flight().HandleChangePassword)
}

func Router() *gin.Engine {
	return router
}
