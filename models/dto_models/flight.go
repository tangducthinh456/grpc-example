package dtomodels

type Flight struct {
	ID          string `json:"id"`
	FlightPlane string `json:"flight_plane" binding:"required" validate:"required"`
	From        string `json:"from" binding:"required" validate:"required"`
	To          string `json:"to" binding:"required" validate:"required"`
	DepartDate  string `json:"depart_date" binding:"required" validate:"required"`
	DepartTime  string `json:"depart_time" binding:"required" validate:"required"`
	Status      string `json:"status" binding:"required" validate:"required"`
}

