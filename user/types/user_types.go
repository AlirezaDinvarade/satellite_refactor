package types

import (
	"satellite/user/models"
)

type CreateUserInput struct {
	NationalID  string `json:"nationalID" validate:"required,len=10"`
	PhoneNumber string `json:"phoneNumber" validate:"required,len=11"`
}

func (c *CreateUserInput) NewUserFromParams() *models.User {
	return &models.User{
		NationalID:  c.NationalID,
		PhoneNumber: c.PhoneNumber,
	}
}
