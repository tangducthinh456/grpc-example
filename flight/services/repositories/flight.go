package repositories

import (
	daomodels "bookingFlight/models/dao_models"
	"context"
	"sync"
)

type flight interface {
	MigrateDataModel(ctx context.Context) error

	CreateFlight(ctx context.Context, model *daomodels.Flight) error
	UpdateFlight(ctx context.Context, id string, model *daomodels.Flight) error
	GetFlightByID(ctx context.Context, id string) (*daomodels.Flight, error)
}

type flightDB struct {
	sync.Mutex
}

var flightDBInstance flight = &flightDB{}

func Flight() flight {
	return flightDBInstance
}

func (c *flightDB) MigrateDataModel(ctx context.Context) (er error) {
	c.Lock()
	defer c.Unlock()

	er = postgresDB.WithContext(ctx).AutoMigrate(&daomodels.Flight{})
	return
}

func (c *flightDB) CreateFlight(ctx context.Context, flight *daomodels.Flight) error {
	c.Lock()
	defer c.Unlock()

	er := postgresDB.WithContext(ctx).Model(&daomodels.Flight{}).Create(flight).Error
	return er

}

func (c *flightDB) GetFlightByID(ctx context.Context, uuid string) (*daomodels.Flight, error) {
	c.Lock()
	defer c.Unlock()

	flight := &daomodels.Flight{}
	err := postgresDB.WithContext(ctx).Model(&daomodels.Flight{}).Where(&daomodels.Flight{UUID: uuid}).First(flight).Error
	return flight, err
}

func (c *flightDB) UpdateFlight(ctx context.Context, uuid string, flight *daomodels.Flight) error {
	c.Lock()
	defer c.Unlock()

	flight.UUID = uuid
	er := postgresDB.WithContext(ctx).Model(&daomodels.Flight{}).Where(&daomodels.Flight{UUID: uuid}).Updates(flight).Error
	return er
}
