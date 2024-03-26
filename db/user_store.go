package db

import (
	"context"
	"errors"

	"github.com/felipemagrassi/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	usersCollection = "users"
	ErrNotFound     = errors.New("user not found")
)

type UserStore interface {
	GetById(context.Context, string) (*types.User, error)
	List(context.Context) ([]*types.User, error)
	Create(context.Context, *types.User) (string, error)
	Update(context.Context, *types.User) error
	Delete(context.Context, string) error
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
		if err == mongo.ErrNoDocuments {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return &foundUser, nil
}

func (s *MongoUserStore) List(ctx context.Context) ([]*types.User, error) {
	var foundUsers []*types.User

	cursor, err := s.coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var user types.User
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		foundUsers = append(foundUsers, &user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return foundUsers, nil
}

func (s *MongoUserStore) Create(ctx context.Context, user *types.User) (string, error) {
	res, err := s.coll.InsertOne(ctx, user)
	if err != nil {
		return "", err
	}

	return FromObjectId(res.InsertedID.(primitive.ObjectID)), nil
}

func (s *MongoUserStore) Update(ctx context.Context, user *types.User) error {
	updateParams := bson.M{
		"$set": bson.M{"email": user.Email},
	}

	oid, err := ToObjectId(user.ID)
	if err != nil {
		return err
	}

	res, err := s.coll.UpdateOne(ctx, bson.M{"_id": oid}, updateParams)
	if err != nil {
		return err
	}

	if res.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}

func (s *MongoUserStore) Delete(ctx context.Context, id string) error {
	oid, err := ToObjectId(id)
	if err != nil {
		return err
	}
	res, err := s.coll.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}
