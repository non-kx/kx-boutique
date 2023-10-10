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

	mongodb, err := getMongoDB(c.Database.Uri, c.Database.DbName)
	if err != nil {
		log.Errorf("Failed connecting to mongodb: %v", err)
		return nil, nil, err
	}

	rediscli, err := getRedisCli(c.Redis.Addr, c.Redis.Pwd, int(c.Redis.Db))

	return &Data{
		db:  mongodb,
		rdb: rediscli,
	}, cleanup, nil
}

func getMongoDB(uri string, name string) (*mongo.Database, error) {
	opt := options.Client().ApplyURI(uri)
	cli, err := mongo.Connect(ctx, opt)
	if err != nil {
		return nil, err
	}

	db := cli.Database(name)

	return db, nil
}

func getRedisCli(addr string, pwd string, db int) (*redis.Client, error) {
	cli := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd, // no password set
		DB:       db,
	})

	return cli, nil
}
