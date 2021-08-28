package daomodels

import (
	"time"
	"gorm.io/gorm"
)

type Flight struct {
	gorm.Model
	UUID          string
	Name          string
	From          string
	To            string
	Date          time.Time
	Status        string
	AvailableSlot int64
}
