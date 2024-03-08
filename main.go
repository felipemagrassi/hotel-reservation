package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/felipemagrassi/hotel-reservation/api"
	"github.com/felipemagrassi/hotel-reservation/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dbrui           = "mongodb://localhost:27017"
	dbname          = "hotel-reservation"
	usersCollection = "users"
)

func main() {
	ctx := context.Background()
	listenAddr := flag.String("listenAddr", ":3000", "The address to listen on")
	flag.Parse()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbrui))
	if err != nil {
		log.Fatal(err)
	}

	user := types.User{
		FirstName: "John",
		LastName:  "Doe",
	}

	coll := client.Database(dbname).Collection(usersCollection)
	_, err = coll.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
	}

	var john types.User
	err = coll.FindOne(ctx, bson.M{"firstName": "John"}).Decode(&john)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(john)

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
