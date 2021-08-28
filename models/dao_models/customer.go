package daomodels

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	UUID        string
	Name        string
	Address     string
	Dob         string
	PhoneNumber string
	Email       string
	Password    string
}
