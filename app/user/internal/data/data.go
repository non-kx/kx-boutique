package data

import (
	"context"
	"fmt"
	"kx-boutique/app/user/internal/conf"

	"kx-boutique/app/user/internal/data/ent"

	"entgo.io/ent/dialect"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	_ "github.com/lib/pq"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	db *ent.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}

	dataSource := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable",
		c.Database.User,
		c.Database.Password,
		c.Database.Url,
		c.Database.Dbname)

	client, err := ent.Open(dialect.Postgres, dataSource)
	if err != nil {
		log.NewHelper(logger).Fatalf("failed opening connection to postgresql: %v", err)
		return nil, nil, err
	}

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.NewHelper(logger).Fatalf("failed creating schema resources: %v", err)
		return nil, nil, err
	}

	return &Data{
		db: client,
	}, cleanup, nil
}
