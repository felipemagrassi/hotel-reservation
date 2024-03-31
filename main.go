package main

import (
	"context"
	"flag"
	"log"

	"github.com/felipemagrassi/hotel-reservation/api"
	"github.com/felipemagrassi/hotel-reservation/db"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	config = fiber.Config{ErrorHandler: errorHandler()}
)

func main() {
	ctx := context.Background()

	listenAddr := flag.String("listenAddr", ":5200", "The address to listen on")
	flag.Parse()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}

	var (
		userStore    = db.NewMongoUserStore(client)
		hotelStore   = db.NewMongoHotelStore(client)
		roomStore    = db.NewMongoRoomStore(client, hotelStore)
		userHandler  = api.NewUserHandler(userStore)
		hotelHandler = api.NewHotelHandler(hotelStore, roomStore)
		app          = fiber.New(config)
		apiv1        = app.Group("/api/v1")
		userv1       = apiv1.Group("/users")
		hotelv1      = apiv1.Group("/hotels")
	)
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		// For more options, see the Config section
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}​\n",
	}))

	userv1.Get("/", userHandler.HandleGetUsers)
	userv1.Get(":id", userHandler.HandleGetUser)
	userv1.Post("/", userHandler.HandleCreateUser)
	userv1.Put(":id", userHandler.HandleEditUser)
	userv1.Delete(":id", userHandler.HandleDeleteUser)

	hotelv1.Get("/", hotelHandler.HandleListHotels)

	app.Listen(*listenAddr)
}

func errorHandler() func(*fiber.Ctx, error) error {
	return func(ctx *fiber.Ctx, err error) error {
		return ctx.JSON(map[string]string{"error": err.Error()})
	}
}
