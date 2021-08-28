package dtomodels

type BookingAssociate struct {
	ID       string   `json:"booking_code"`
	Date     string   `json:"booking_date"`
	Customer Customer `json:"customer" `
	Flight   Flight   `json:"flight"`
}