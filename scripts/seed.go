package main

import (
	"context"
	"log"

	"github.com/felipemagrassi/hotel-reservation/db"
	"github.com/felipemagrassi/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}

	seedHotel(ctx, client, "Hotel A", "Location A")
	seedHotel(ctx, client, "Hotel B", "Location B")
	seedHotel(ctx, client, "Hotel C", "Location C")
}

func seedHotel(ctx context.Context, client *mongo.Client, name, location string) {
	hotelStore := db.NewMongoHotelStore(client)
	roomStore := db.NewMongoRoomStore(client, hotelStore)

	hotel := types.Hotel{
		Name:     name,
		Location: location,
		Rooms:    []string{},
	}

	insertedHotelID, err := hotelStore.InsertHotel(ctx, &hotel)
	if err != nil {
		log.Fatal(err)
	}

	rooms := []types.Room{
		{
			Type:      types.SingleRoomType,
			HotelID:   insertedHotelID,
			BasePrice: 99.9,
		},
		{

			Type:      types.DoubleRoomType,
			HotelID:   insertedHotelID,
			BasePrice: 199.9,
		},
	}

	for _, room := range rooms {
		_, err := roomStore.InsertRoom(ctx, &room)
		if err != nil {
			log.Fatal(err)
		}
	}

}
