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
	users, err := h.userStore.List(c.Context())

	if err != nil {
		return err
	}

	return c.JSON(users)
}

func (h *UserHandler) HandleDeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.userStore.Delete(c.Context(), id); err != nil {
		return err
	}

	return c.JSON(map[string]string{"deleted": id})
}

func (h *UserHandler) HandleEditUser(c *fiber.Ctx) error {
	var dto types.UserDTO
	id := c.Params("id")

	if err := c.BodyParser(&dto); err != nil {
		return err
	}

	if errs, ok := dto.Validate(); !ok {
		return c.Status(fiber.StatusBadRequest).JSON(errs)
	}

	user, err := types.NewUserUpdateDTO(dto)
	if err != nil {
		return err
	}

	user.ID = id

	err = h.userStore.Update(c.Context(), user)
	if err != nil {
		return err
	}

	return c.JSON(map[string]string{"updated": id})
}

func (h *UserHandler) HandleCreateUser(c *fiber.Ctx) error {
	var dto types.UserDTO

	if err := c.BodyParser(&dto); err != nil {
		return err
	}

	errs, ok := dto.Validate()
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(errs)
	}

	user, err := types.NewUserCreateDTO(dto)
	if err != nil {
		return err
	}

	id, err := h.userStore.Create(c.Context(), user)
	if err != nil {
		return err
	}

	user.ID = id

	return c.JSON(user)
}
