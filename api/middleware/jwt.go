package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func JWTAuthorization(c *fiber.Ctx) error {
	token, ok := c.GetReqHeaders()["X-Api-Token"]
	if !ok {
		fmt.Println("JWT authorization failed")
		return fmt.Errorf("Unauthorized")
	}

	fmt.Println("token:", token)
	fmt.Println("JWT authorization success")

	return nil
}
