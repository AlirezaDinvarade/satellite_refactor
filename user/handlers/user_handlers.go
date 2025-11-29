package handlers

import (
	"encoding/json"
	"net/http"
	"satellite/user/stores"
	"satellite/user/types"
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

func (h *UserHandler) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	var userParams types.CreateUserInput
	if err := json.NewDecoder(r.Body).Decode(&userParams); err != nil {
		ErrorInvalidData(w)
		return
	}

	if err := validate.Struct(userParams); err != nil {
		ErrorInvalidData(w)
		return
	}

	user := userParams.NewUserFromParams()
	user, err := h.store.User.InsertUser(r.Context(), user)
	if err != nil {
		ErrorInternalServerError(w)
		return
	}

}
