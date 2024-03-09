package db

import (
	"context"

	"github.com/felipemagrassi/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	usersCollection = "users"
)

type UserStore interface {
	GetById(ctx context.Context, id string) (*types.User, error)
}

func NewMongoUserStore(client *mongo.Client) *MongoUserStore {
	return &MongoUserStore{
		client: client,
		coll: client.Database(
			GetDBName(),
		).Collection(usersCollection),
	}
}

type MongoUserStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func (s *MongoUserStore) GetById(ctx context.Context, id string) (*types.User, error) {
	var foundUser types.User
	oid, err := ToObjectId(id)
	if err != nil {
		return nil, err
	}

	err = s.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&foundUser)
	if err != nil {
		return nil, err
	}

	return &foundUser, nil
}
