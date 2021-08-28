package repositories

import (
	daomodels "bookingFlight/models/dao_models"
	"context"

	// "gorm.io/gorm"
	"sync"
)

type booking interface {
	MigrateDataModel(ctx context.Context) error

	CreateBooking(ctx context.Context, model *daomodels.Booking) error
	GetListBookingByCustomer(ctx context.Context, id string) ([]daomodels.Booking, error)
	GetBookingByID(ctx context.Context, uuid string) (*daomodels.Booking, error)
	DeleteBooking(ctx context.Context, id string) error
}

type bookingDB struct {
	sync.Mutex
}

var bookingDBInstance booking = &bookingDB{}

func Booking() booking {
	return bookingDBInstance
}

func (c *bookingDB) MigrateDataModel(ctx context.Context) (er error) {
	c.Lock()
	defer c.Unlock()

	er = postgresDB.WithContext(ctx).AutoMigrate(&daomodels.Booking{})
	return
}

func (c *bookingDB) CreateBooking(ctx context.Context, booking *daomodels.Booking) error {
	c.Lock()
	defer c.Unlock()

	er := postgresDB.WithContext(ctx).Model(&daomodels.Booking{}).Create(booking).Error
	return er
}

func (c *bookingDB) GetListBookingByCustomer(ctx context.Context, customer_id string) ([]daomodels.Booking, error) {
	c.Lock()
	defer c.Unlock()

	var listbook []daomodels.Booking

	err := postgresDB.WithContext(ctx).Model(&daomodels.Booking{}).Where(&daomodels.Booking{CustomerId: customer_id}).Find(&listbook).Error
	return listbook, err
}

func (c *bookingDB) GetBookingByID(ctx context.Context, uuid string) (*daomodels.Booking, error) {
	c.Lock()
	defer c.Unlock()

	booking := &daomodels.Booking{}
	err := postgresDB.WithContext(ctx).Model(&daomodels.Booking{}).Where(&daomodels.Booking{UUID: uuid}).First(booking).Error
	return booking, err
}

func (c *bookingDB) DeleteBooking(ctx context.Context, uuid string) error {
	c.Lock()
	defer c.Unlock()

	er := postgresDB.WithContext(ctx).Model(&daomodels.Booking{}).Where(&daomodels.Booking{UUID: uuid}).Delete(&daomodels.Booking{}).Error
	return er
}
