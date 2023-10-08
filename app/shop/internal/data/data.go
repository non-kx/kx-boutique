package data

import (
	"context"
	"kx-boutique/app/shop/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewProductRepo)
var collection *mongo.Collection
var ctx = context.TODO()

// Data .
type Data struct {
	// TODO wrapped database client
	db  *mongo.Database
	rdb *redis.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}

	clientOptions := options.Client().ApplyURI(c.Database.Uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Errorf("Failed connecting to mongodb: %v", err)
		return nil, nil, err
	}

	db := client.Database(c.Database.DbName)

	return &Data{
		db,
		nil,
	}, cleanup, nil
}
