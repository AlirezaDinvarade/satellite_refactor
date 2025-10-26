package types

import (
	"satellite/user/models"
)

type CreateUserParams struct {
	NationalID  string `json:"nationalID" validate:"required"`
	PhoneNumber string `json:"phoneNumber" validate:"required"`
}

func (c *CreateUserParams) NewUserFromParams() *models.User {
	return &models.User{
		NationalID:  c.NationalID,
		PhoneNumber: c.PhoneNumber,
	}
}
