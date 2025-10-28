package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	if apiError, ok := err.(Error); ok {
		return c.Status(apiError.Code).JSON(apiError)
	}

	apiError := NewError(http.StatusInternalServerError, err.Error())
	return c.Status(apiError.Code).JSON(apiError)

}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e Error) Error() string {
	return e.Message
}

func NewError(code int, message string) Error {
	return Error{
		Code:    code,
		Message: message,
	}
}

func ErrorInvalidData() Error {
	return Error{
		Code:    http.StatusBadRequest,
		Message: "input values are not valid",
	}
}

func ErrorInternalServerError() Error {
	return Error{
		Code:    http.StatusInternalServerError,
		Message: "Internal server error",
	}
}

func ErrorActiveOTP() Error {
	return Error{
		Code:    http.StatusBadRequest,
		Message: "you have active otp",
	}
}

func ErrorExpireOTP() Error {
	return Error{
		Code:    http.StatusBadRequest,
		Message: "Your OTP code has been expired",
	}
}

func ErrorMissMatchOTP() Error {
	return Error{
		Code:    http.StatusBadRequest,
		Message: "Your OTP code is miss match",
	}
}
