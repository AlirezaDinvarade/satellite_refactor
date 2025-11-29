package handlers

import (
	"encoding/json"
	"net/http"
)

func WriteJson(rw http.ResponseWriter, Code int, v interface{}) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(Code)
	json.NewEncoder(rw).Encode(v)
}

func ErrorInvalidData(rw http.ResponseWriter) {
	WriteJson(rw, http.StatusBadRequest, map[string]string{"detail": "input values are not valid"})
}

func ErrorInternalServerError(rw http.ResponseWriter) {
	WriteJson(rw, http.StatusInternalServerError, map[string]string{"detail": "Internal server error"})
}

func ErrorActiveOTP(rw http.ResponseWriter) {
	WriteJson(rw, http.StatusBadRequest, map[string]string{"detail": "you have active otp"})
}

func ErrorExpireOTP(rw http.ResponseWriter) {
	WriteJson(rw, http.StatusBadRequest, map[string]string{"detail": "Your OTP code has been expired"})
}

func ErrorMissMatchOTP(rw http.ResponseWriter) {
	WriteJson(rw, http.StatusBadRequest, map[string]string{"detail": "Your OTP code is miss match"})
}

func ErrorMissMatchPasswords(rw http.ResponseWriter) {
	WriteJson(rw, http.StatusBadRequest, map[string]string{"detail": "Your passwords is not match together"})
}
