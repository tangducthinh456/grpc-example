package config

import (
	"bookingFlight/log"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type PostgresqlConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"db"`
	User     string `yaml:"user"`
	Password string `yaml:"pwd"`
}

type ServerConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type ServiceConfig struct {
	API  ServerConfig `yaml:"api_server"`
	GRPC ServerConfig `yaml:"grpc_server"`
}

type config struct {
	Customer ServiceConfig    `yaml:"customer_service"`
	Flight   ServiceConfig    `yaml:"flight_service"`
	Booking  ServiceConfig    `yaml:"booking_service"`
	Postgres PostgresqlConfig `yaml:"postgresql"`
}

var (
	conf            config
	customerService ServiceConfig
	flight          ServiceConfig
	booking         ServiceConfig
	psgConfig       PostgresqlConfig
)

func Init() {
	log.Info("Start initialize config")
	initConfig, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Error("Initial config fail : " + err.Error())
		panic(err)
	}

	err = yaml.Unmarshal(initConfig, &conf)
	if err != nil {
		log.Error("Initial config failed : " + err.Error())
	}

	customerService = conf.Customer

	psgConfig = conf.Postgres

	booking = conf.Booking

	flight = conf.Flight

}

func CustomerConfig() ServiceConfig {
	return customerService
}

func PsgConfig() PostgresqlConfig {
	return psgConfig
}

func BookingConfig() ServiceConfig {
	return booking
}

func FlightConfig() ServiceConfig {
	return flight
}
