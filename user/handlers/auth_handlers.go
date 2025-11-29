package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"satellite/user/models"
	"satellite/user/stores"
	"satellite/user/types"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type AuthHandler struct {
	DatabaseStore *gorm.DB
	CacheStore    stores.RedisClient
}

func NewAuthHandler(authStore *gorm.DB, cacheStore stores.RedisClient) *AuthHandler {
	return &AuthHandler{
		DatabaseStore: authStore,
		CacheStore:    cacheStore,
	}
}

func (h *AuthHandler) SendOTPHandler(w http.ResponseWriter, r *http.Request) {
	var params types.SendOTPInput
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		ErrorInvalidData(w)
		return
	}

	if err := validate.Struct(params); err != nil {
		ErrorInvalidData(w)
		return
	}

	key := fmt.Sprintf("otp:%s", params.PhoneNumber)
	value, err := h.CacheStore.Get(r.Context(), key)
	if value != "" {
		ErrorActiveOTP(w)
		return
	}

	var user *models.User
	if err := h.DatabaseStore.Where("phone_number = ?", params.PhoneNumber).First(&user).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		ErrorInternalServerError(w)
		return
	}

	if user != nil && user.Password != "" {
		WriteJson(w, http.StatusOK, types.SendOTPResponse{
			Otp:         nil,
			HasPassword: true,
		})
		return
	}

	otp := strconv.Itoa(rand.Intn(90000) + 10000)
	if err = h.CacheStore.SetEx(r.Context(), key, []byte(otp), time.Minute*2); err != nil {
		ErrorInternalServerError(w)
		return
	}

	WriteJson(w, http.StatusOK, types.SendOTPResponse{
		Otp:         &otp,
		HasPassword: bool(user != nil && user.Password != ""),
	})

}

func (h *AuthHandler) LoginOTPHandler(w http.ResponseWriter, r *http.Request) {
	var params types.OTPLoginInput
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil || validate.Struct(params) != nil {
		ErrorInvalidData(w)
		return
	}

	key := fmt.Sprintf("otp:%s", params.PhoneNumber)
	OTPStored, err := h.CacheStore.Get(r.Context(), key)
	if OTPStored == "" || err != nil {
		ErrorExpireOTP(w)
		return
	}
	if OTPStored != params.Otp {
		ErrorMissMatchOTP(w)
		return
	}

	var user *models.User
	if err = h.DatabaseStore.FirstOrCreate(&user, models.User{PhoneNumber: params.PhoneNumber}).Error; err != nil {
		ErrorInternalServerError(w)
		return
	}

	session := types.RedisSessionData{
		PhoneNumber: user.PhoneNumber,
		AccessLevel: string(user.AccessLevel),
	}
	residValue, err := json.Marshal(session)
	token := "some-random-token"
	activeSessionTTL := 24 * time.Hour
	err = h.CacheStore.SetEx(r.Context(), token, residValue, activeSessionTTL)
	if err != nil {
		ErrorInternalServerError(w)
		return
	}

	if err = h.CacheStore.Del(r.Context(), key); err != nil {
		ErrorInternalServerError(w)
	}

	WriteJson(w, http.StatusOK, types.LoginResponse{
		Token: token,
		User:  *user,
	})
}
