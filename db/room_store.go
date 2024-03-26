package db

import (
	"context"
	"errors"

	"github.com/felipemagrassi/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RoomStore interface {
	InsertRoom(context.Context, *types.Room) (string, error)
}

type MongoRoomStore struct {
	client     *mongo.Client
	collection *mongo.Collection

	HotelStore
}

func NewMongoRoomStore(client *mongo.Client, hotelStore HotelStore) *MongoRoomStore {
	return &MongoRoomStore{
		client:     client,
		collection: client.Database(DBNAME).Collection("rooms"),
		HotelStore: hotelStore,
	}
}

func (s *MongoRoomStore) InsertRoom(ctx context.Context, room *types.Room) (string, error) {
	res, err := s.collection.InsertOne(ctx, room)
	if err != nil {
		panic(err)
	}

	if res.InsertedID == nil {
		return "", errors.New("failed to insert room")
	}

	room.ID = FromObjectId(res.InsertedID.(primitive.ObjectID))

	if err := s.HotelStore.InsertRoom(ctx, room); err != nil {
		return "", err
	}

	return FromObjectId(res.InsertedID.(primitive.ObjectID)), nil
}
