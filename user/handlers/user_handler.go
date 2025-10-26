package handlers

import (
	"net/http"
	"satellite/user/types"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserHandler struct {
	store *gorm.DB
}

func NewUserHandler(userStore *gorm.DB) *UserHandler {
	return &UserHandler{
		store: userStore,
	}
}

var validate = validator.New()

func (h *UserHandler) HandleCreateUser(c *fiber.Ctx) error {
	var userParams types.CreateUserParams
	if err := c.BodyParser(&userParams); err != nil {
		return err
	}

	if err := validate.Struct(userParams); err != nil {
		return err
	}

	user := userParams.NewUserFromParams()
	if err := h.store.Create(&user).Error; err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(userParams)
}
