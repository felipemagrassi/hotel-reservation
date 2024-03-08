package api

import (
	"github.com/felipemagrassi/hotel-reservation/types"
	"github.com/gofiber/fiber/v2"
)

func HandleGetUser(c *fiber.Ctx) error {
	user := types.User{
		FirstName: "John",
		LastName:  "Doe",
	}
	return c.JSON(user)
}

func HandleGetUsers(c *fiber.Ctx) error {
	users := []types.User{
		{FirstName: "John", LastName: "Doe"},
		{FirstName: "Jane", LastName: "Doe"},
	}

	return c.JSON(users)
}
