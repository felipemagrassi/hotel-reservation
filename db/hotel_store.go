package db

import (
	"context"

	"github.com/felipemagrassi/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type HotelStore interface {
	InsertHotel(context.Context, *types.Hotel) (string, error)
	Update(ctx context.Context, hotel *types.Hotel) error
}

type MongoHotelStore struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewMongoHotelStore(client *mongo.Client) *MongoHotelStore {
	return &MongoHotelStore{
		client:     client,
		collection: client.Database(DBNAME).Collection("hotels"),
	}
}

func (s *MongoHotelStore) InsertHotel(ctx context.Context, hotel *types.Hotel) (string, error) {
	res, err := s.collection.InsertOne(ctx, hotel)
	if err != nil {
		panic(err)
	}

	return FromObjectId(res.InsertedID.(primitive.ObjectID)), nil
}

func (s *MongoHotelStore) UpdateRoom(ctx context.Context, hotel *types.Hotel) error {
	oid, err := ToObjectId(hotel.ID)
	if err != nil {
		return err
	}

	filter := primitive.M{"_id": oid}
	params := primitive.M{"$push": bson.M{"rooms": hotel.Rooms}}

}
