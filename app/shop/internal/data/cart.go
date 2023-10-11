package data

import (
	"context"
	"encoding/json"
	"time"

	"kx-boutique/app/shop/internal/biz"
	"kx-boutique/app/shop/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
)

const (
	CART_KEY_PREFIX = "cart-"
	DEFUALT_TTL     = 168 * time.Hour
)

type cartRepo struct {
	rdb *redis.Client
	log *log.Helper
}

func NewCartRepo(c *conf.Data, data *Data, logger log.Logger) biz.CartRepo {
	return &cartRepo{
		rdb: data.rdb,
		log: log.NewHelper(logger),
	}
}

func (r *cartRepo) GetUserCart(ctx context.Context, userId string) (*biz.Cart, error) {
	raw, err := r.rdb.Get(ctx, getCartKey(userId)).Result()
	if err != nil {
		if err == redis.Nil {
			return &biz.Cart{UserId: userId}, nil
		}
		return nil, err
	}

	var cart biz.Cart
	err = json.Unmarshal([]byte(raw), &cart)
	if err != nil {
		return nil, err
	}

	return &cart, nil
}

func (r *cartRepo) UpdateUserCart(ctx context.Context, userId string, cart *biz.Cart) error {
	raw, err := json.Marshal(*cart)
	if err != nil {
		return err
	}

	r.log.WithContext(ctx).Debug("Update user[%v] cart: ", userId, string(raw))

	err = r.rdb.Set(ctx, getCartKey(userId), raw, DEFUALT_TTL).Err()
	if err != nil {
		return err
	}

	return nil
}

func getCartKey(userId string) string {
	return CART_KEY_PREFIX + userId
}
