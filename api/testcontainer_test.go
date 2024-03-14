package api

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/felipemagrassi/hotel-reservation/db"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mongodb"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	testDatabase *TestDatabase
)

type TestDatabase struct {
	Container testcontainers.Container
	Store     db.UserStore
}

func (t *TestDatabase) Close() {
	if err := t.Container.Terminate(context.Background()); err != nil {
		log.Fatalf("Could not terminate container: %s", err)
	}
}

func setup() (*TestDatabase, error) {
	ctx := context.Background()

	mongodbContainer, err := mongodb.RunContainer(
		ctx,
		testcontainers.WithImage("mongo:6"),
		testcontainers.WithWaitStrategy(wait.ForLog("Waiting for connections")),
	)
	if err != nil {
		log.Fatalf("Could not start mongodb container: %s", err)
	}

	endpoint, err := mongodbContainer.ConnectionString(ctx)
	if err != nil {
		log.Fatalf("Could not get mongodb container connection string: %s", err)
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(endpoint))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	store := db.NewMongoUserStore(client)

	return &TestDatabase{Container: mongodbContainer, Store: store}, nil
}

func TestMain(m *testing.M) {
	testDb, err := setup()
	testDatabase = testDb

	if err != nil {
		log.Fatalf("Could not setup test database: %s", err)
	}

	if testDatabase == nil {
		log.Fatal("Test database not set up")
	}

	log.Println("Test database setup complete")

	exitCode := m.Run()

	if testDb != nil {
		testDb.Close()
	}

	os.Exit(exitCode)
}
