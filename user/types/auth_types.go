package types

import "satellite/user/models"

type SendOTPInput struct {
	PhoneNumber string `json:"phoneNumber" validate:"required,len=11,startswith09"`
}

type SendOTPResponse struct {
	Otp         *string `json:"otp"`
	HasPassword bool    `json:"has_password"`
}

type OTPLoginInput struct {
	Otp         string `json:"otp"`
	PhoneNumber string `json:"phoneNumber" validate:"required,len=11,startswith09"`
}

type RedisSessionData struct {
	PhoneNumber string `json:"phone_number"`
	AccessLevel string `json:"access_level"`
}

type LoginResponse struct {
	Token string      `json:"token"`
	User  models.User `json:"user"`
}

type SetPasswordInput struct {
	Password        string `json:"password"`
	ConfirmPassword string `json:"confim_password"`
}
