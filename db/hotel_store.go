package db

import (
	"context"
	"errors"

	"github.com/felipemagrassi/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type HotelStore interface {
	InsertHotel(context.Context, *types.Hotel) (string, error)
	InsertRoom(ctx context.Context, room *types.Room) error
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

func (s *MongoHotelStore) GetHoteL(ctx context.Context, id string) (*types.Hotel, error) {
	oid, err := ToObjectId(id)
	if err != nil {
		return nil, err
	}

	var hotel types.Hotel

	err = s.collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&hotel)
	if err != nil {
		return nil, err
	}

	return &hotel, nil
}

func (s *MongoHotelStore) ListHotels(ctx context.Context) ([]*types.Hotel, error) {
	cursor, err := s.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var hotels []*types.Hotel
	if err = cursor.All(ctx, &hotels); err != nil {
		return nil, err
	}
	return hotels, nil
}

func (s *MongoHotelStore) InsertRoom(ctx context.Context, room *types.Room) error {
	if room.HotelID == "" {
		return errors.New("Room hotelID is required")
	}

	hotelOid, err := ToObjectId(room.HotelID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": hotelOid}
	update := bson.M{"$push": bson.M{"rooms": room.ID}}

	res, err := s.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if res.ModifiedCount == 0 {
		return errors.New("hotel not found")
	}

	return nil

}
