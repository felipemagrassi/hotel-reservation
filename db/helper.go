package db

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	DBNAME = "hotel-reservation"
	TESTDB = "test-hotel-reservation"
	DBURI  = "mongodb://localhost:27017"
)

func GetDBName() string {
	return DBNAME
}

func ToObjectId(id string) (primitive.ObjectID, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return oid, nil
}

func FromObjectId(oid primitive.ObjectID) string {
	return oid.Hex()
}
