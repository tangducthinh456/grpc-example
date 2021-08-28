package repositories

import (
	"context"
	daomodels "bookingFlight/models/dao_models"
	"sync"
	"gorm.io/gorm"
)

type customer interface {
	MigrateDataModel(ctx context.Context) error

	CreateCustomer(ctx context.Context, model *daomodels.Customer) error
	GetCustomerByID(ctx context.Context, id string) (*daomodels.Customer, error)
	UpdateCustomer(ctx context.Context, id string, model *daomodels.Customer) error
	DeleteCustomer(ctx context.Context, id string) error
}

type customerDB struct {
	sync.Mutex
}

var customerDBInstance customer = &customerDB{}

func Customer() customer {
	return customerDBInstance
}

func (c *customerDB) MigrateDataModel(ctx context.Context) (er error) {
	c.Lock()
	defer c.Unlock()
	
	er = postgresDB.WithContext(ctx).AutoMigrate(&daomodels.Customer{})
	return
}

func (c *customerDB) CreateCustomer(ctx context.Context, customer *daomodels.Customer) error {
	c.Lock()
	defer c.Unlock()

	return postgresDB.WithContext(ctx).Transaction(func(tx *gorm.DB) (er error) {
		er = tx.WithContext(ctx).Model(&daomodels.Customer{}).Create(customer).Error
		return er
	})
}

func (c *customerDB) GetCustomerByID(ctx context.Context, uuid string) (*daomodels.Customer, error) {
	c.Lock()
	defer c.Unlock()

	customer := &daomodels.Customer{}
	err := postgresDB.WithContext(ctx).Model(&daomodels.Customer{}).Where(&daomodels.Customer{UUID: uuid}).First(customer).Error
	return customer, err
}

func (c *customerDB) UpdateCustomer(ctx context.Context, uuid string, customer *daomodels.Customer) error {
	c.Lock()
	defer c.Unlock()
	
	customer.UUID = uuid
	er := postgresDB.WithContext(ctx).Model(&daomodels.Customer{}).Where(&daomodels.Customer{UUID: uuid}).Updates(customer).Error
    return er
}

func (c *customerDB) DeleteCustomer(ctx context.Context, uuid string) error {
	c.Lock()
	defer c.Unlock()

	err := postgresDB.WithContext(ctx).Model(&daomodels.Customer{}).Where(&daomodels.Customer{UUID: uuid}).Delete(&daomodels.Customer{}).Error
    return err
}
