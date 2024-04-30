package mongo

import (
	"context"
	"github.com/romeros69/basket/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	Db *mongo.Database
}

func New(cfg *config.Config) (*Mongo, error) {
	ctx := context.Background()
	cOpts := options.Client().ApplyURI(cfg.MongoURL)
	mClient, err := mongo.Connect(ctx, cOpts)
	if err != nil {
		return nil, err
	}
	err = mClient.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}
	return &Mongo{
		Db: mClient.Database(cfg.MongoDB),
	}, nil
}
