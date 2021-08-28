package repositories

import (
	"bookingFlight/config"
	"bookingFlight/log"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	postgresDB *gorm.DB
)

func InitDatabaseConnection() (err error) {
	log.Info("Init database postgres connection")
	conf := config.PsgConfig()
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%s host=%s sslmode=disable",
		conf.User, conf.Password, conf.Database,
		conf.Port, conf.Host)
	postgresDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
	return
}
