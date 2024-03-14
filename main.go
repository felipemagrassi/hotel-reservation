package main

import (
	"context"
	"flag"
	"log"

	"github.com/felipemagrassi/hotel-reservation/api"
	"github.com/felipemagrassi/hotel-reservation/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx := context.Background()

	listenAddr := flag.String("listenAddr", ":3000", "The address to listen on")
	flag.Parse()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}

	mongoStore := db.NewMongoUserStore(client)
	userHandler := api.NewUserHandler(mongoStore)

	app := fiber.New(
		fiber.Config{
			ErrorHandler: func(ctx *fiber.Ctx, err error) error {
				return ctx.JSON(map[string]string{"error": err.Error()})
			},
		},
	)

	appApi := app.Group("/api")
	appApiV1 := appApi.Group("/v1")

	appApiV1User := appApiV1.Group("/user")
	appApiV1User.Get("/", userHandler.HandleGetUsers)
	appApiV1User.Get(":id", userHandler.HandleGetUser)
	appApiV1User.Post("/", userHandler.HandleCreateUser)
	appApiV1User.Put(":id", userHandler.HandleEditUser)
	appApiV1User.Delete(":id", userHandler.HandleDeleteUser)

	app.Listen(*listenAddr)
}

func handleFoo(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"hello": "world"})
}
