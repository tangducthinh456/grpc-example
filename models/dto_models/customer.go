package dtomodels

type Customer struct {
	ID      string `json:"id"`
	Name    string `json:"name" binding:"required" validate:"required"`
	Dob     string `json:"date_of_birth" binding:"required" validate:"required"`
	Address string `json:"address" binding:"required" validate:"required"`
	Email   string `json:"email" binding:"required" validate:"required"`
	Phone   string `json:"phone" binding:"required" validate:"required"`
}

type RequestChangePassword struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password" validate:"required"`
}