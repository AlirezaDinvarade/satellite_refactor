package handlers

import (
	"gorm.io/gorm"
)

type AuthHandler struct {
	store *gorm.DB
}

func NewAuthHandler(authStore *gorm.DB) *AuthHandler {
	return &AuthHandler{
		store: authStore,
	}
}
