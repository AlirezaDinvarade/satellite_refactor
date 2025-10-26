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
	var userParams types.CreateUserInput
	if err := c.BodyParser(&userParams); err != nil {
		return err
	}

	if err := validate.Struct(userParams); err != nil {
		return ErrorInvalidData()
	}

	user := userParams.NewUserFromParams()
	if err := h.store.Create(&user).Error; err != nil {
		return ErrorInternalServerError()
	}

	return c.Status(http.StatusOK).JSON(user)
}
