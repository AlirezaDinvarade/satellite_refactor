package handlers

import (
	"net/http"
	"satellite/user/stores"
	"satellite/user/types"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	store *stores.Store
}

func NewUserHandler(store *stores.Store) *UserHandler {
	return &UserHandler{
		store: store,
	}
}

var validate = types.NewValidator()

func (h *UserHandler) HandleCreateUser(c *fiber.Ctx) error {
	var userParams types.CreateUserInput
	if err := c.BodyParser(&userParams); err != nil {
		return ErrorInvalidData()
	}

	if err := validate.Struct(userParams); err != nil {
		return ErrorInvalidData()
	}

	user := userParams.NewUserFromParams()
	user, err := h.store.User.InsertUser(c.Context(), user)
	if err != nil {
		return ErrorInternalServerError()
	}

	return c.Status(http.StatusOK).JSON(user)
}
