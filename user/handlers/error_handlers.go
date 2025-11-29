package handlers

import (
	"encoding/json"
	"net/http"
)

func WriteJson(rw http.ResponseWriter, Code int, v interface{}) {
	rw.WriteHeader(Code)
	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(v)
}

func ErrorInvalidData(rw http.ResponseWriter) {
	WriteJson(rw, http.StatusBadRequest, "input values are not valid")
}

func ErrorInternalServerError(rw http.ResponseWriter) {
	WriteJson(rw, http.StatusInternalServerError, "Internal server error")
}

func ErrorActiveOTP(rw http.ResponseWriter) {
	WriteJson(rw, http.StatusBadRequest, "you have active otp")
}

func ErrorExpireOTP(rw http.ResponseWriter) {
	WriteJson(rw, http.StatusBadRequest, "Your OTP code has been expired")
}

func ErrorMissMatchOTP(rw http.ResponseWriter) {
	WriteJson(rw, http.StatusBadRequest, "Your OTP code is miss match")
}

func ErrorMissMatchPasswords(rw http.ResponseWriter) {
	WriteJson(rw, http.StatusBadRequest, "Your passwords is not match together")
}
