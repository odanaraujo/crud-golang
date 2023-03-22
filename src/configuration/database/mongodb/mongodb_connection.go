package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

var (
	MONGODB_URL      = "MONGODB_URL"
	MONGODB_DATABASE = "MONGODB_DATABASE"
)

func NewMongodbConnection(ctx context.Context) (*mongo.Database, error) {
	mongoURL := os.Getenv(MONGODB_URL)
	mongodbDatabase := os.Getenv(MONGODB_DATABASE)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))

	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client.Database(mongodbDatabase), nil

}
