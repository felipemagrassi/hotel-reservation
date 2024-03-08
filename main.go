package main

import (
	"flag"

	"github.com/felipemagrassi/hotel-reservation/api"
	"github.com/gofiber/fiber/v2"
)

func main() {
	listenAddr := flag.String("listenAddr", ":3000", "The address to listen on")
	flag.Parse()

	app := fiber.New()

	appApi := app.Group("/api")
	appApiV1 := appApi.Group("/v1")

	appApiV1User := appApiV1.Group("/user")
	appApiV1User.Get("/", api.HandleGetUsers)
	appApiV1User.Get(":id", api.HandleGetUser)

	app.Listen(*listenAddr)
}

func handleFoo(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"hello": "world"})
}
