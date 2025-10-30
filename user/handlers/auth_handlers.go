package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"satellite/user/models"
	"satellite/user/types"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AuthHandler struct {
	DatabaseStore *gorm.DB
	CacheStore    models.RedisClient
}

func NewAuthHandler(authStore *gorm.DB, cacheStore models.RedisClient) *AuthHandler {
	return &AuthHandler{
		DatabaseStore: authStore,
		CacheStore:    cacheStore,
	}
}

func (h *AuthHandler) SendOTPHandler(c *fiber.Ctx) error {
	var params types.SendOTPInput
	if err := c.BodyParser(&params); err != nil {
		return ErrorInvalidData()
	}

	if err := validate.Struct(params); err != nil {
		return ErrorInvalidData()
	}

	key := fmt.Sprintf("otp:%s", params.PhoneNumber)
	value, err := h.CacheStore.Get(c.Context(), key)
	if value != "" {
		return ErrorActiveOTP()
	}

	var user *models.User
	if err := h.DatabaseStore.Where("phone_number = ?", params.PhoneNumber).First(&user).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrorInternalServerError()
	}

	if user != nil && user.Password != "" {
		return c.Status(http.StatusOK).JSON(types.SendOTPResponse{
			Otp:         nil,
			HasPassword: true,
		})
	}

	otp := strconv.Itoa(rand.Intn(90000) + 10000)
	if err = h.CacheStore.SetEx(c.Context(), key, []byte(otp), time.Minute*2); err != nil {
		return ErrorInternalServerError()
	}

	return c.Status(http.StatusOK).JSON(types.SendOTPResponse{
		Otp:         &otp,
		HasPassword: bool(user != nil && user.Password != ""),
	})

}

func (h *AuthHandler) LoginOTPHandler(c *fiber.Ctx) error {
	var params types.OTPLoginInput
	if err := c.BodyParser(&params); err != nil || validate.Struct(params) != nil {
		return ErrorInvalidData()
	}

	key := fmt.Sprintf("otp:%s", params.PhoneNumber)
	OTPStored, err := h.CacheStore.Get(c.Context(), key)
	if OTPStored == "" || err != nil {
		return ErrorExpireOTP()
	}
	if OTPStored != params.Otp {
		return ErrorMissMatchOTP()
	}

	var user *models.User
	if err = h.DatabaseStore.FirstOrCreate(&user, models.User{PhoneNumber: params.PhoneNumber}).Error; err != nil {
		return ErrorInternalServerError()
	}

	session := types.RedisSessionData{
		PhoneNumber: user.PhoneNumber,
		AccessLevel: string(user.AccessLevel),
	}
	residValue, err := json.Marshal(session)
	token := "some-random-token"
	activeSessionTTL := 24 * time.Hour
	err = h.CacheStore.SetEx(c.Context(), token, residValue, activeSessionTTL).Err()
	if err != nil {
		return ErrorInternalServerError()
	}

	if err = h.CacheStore.Del(c.Context(), key); err != nil {
		return ErrorInternalServerError()
	}

	return c.Status(http.StatusOK).JSON(types.LoginResponse{
		Token: token,
		User:  *user,
	})
}

// func (h *AuthHandler) SetPasswordHandler(c *fiber.Ctx) error {
// 	var params types.SetPasswordInput
// 	if err := c.BodyParser(&params); err != nil {
// 		return ErrorInvalidData()
// 	}

// 	if params.Password != params.ConfirmPassword {
// 		return ErrorMissMatchPasswords()
// 	}

// }
