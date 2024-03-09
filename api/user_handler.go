package api

import (
	"github.com/felipemagrassi/hotel-reservation/db"
	"github.com/felipemagrassi/hotel-reservation/types"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (h *UserHandler) HandleGetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := h.userStore.GetById(c.Context(), id)

	if err != nil {
		return err
	}

	return c.JSON(user)
}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	users := []types.User{
		{FirstName: "John", LastName: "Doe"},
		{FirstName: "Jane", LastName: "Doe"},
	}

	return c.JSON(users)
}
