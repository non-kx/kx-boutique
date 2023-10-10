package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

var (
// ErrProductNotFound is product not found.
// ErrProductNotFound = errors.NotFound(v1.ErrorReason_PRODUCT_NOT_FOUND.String(), "product not found")
)

type Cart struct {
	UserId string `json:"user_id"`
	Items  []Item `json:"items"`
}

type Item struct {
	ProductId string `json:"product_id"`
	Qty       int64  `json:"qty"`
}

type CartRepo interface {
	GetUserCart(context.Context, string) (*Cart, error)
	UpdateUserCart(context.Context, string, *Cart) error
}

type CartUsecase struct {
	repo CartRepo
	log  *log.Helper
}

func NewCartUsecase(repo CartRepo, logger log.Logger) *CartUsecase {
	return &CartUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *CartUsecase) GetUserCart(ctx context.Context, userId string) (*Cart, error) {
	uc.log.WithContext(ctx).Infof("Get user[%v] cart: %v", userId)

	cart, err := uc.repo.GetUserCart(ctx, userId)
	if err != nil {
		return nil, err
	}

	return cart, nil
}

func (uc *CartUsecase) AddToCart(ctx context.Context, i Item, userId string) error {
	uc.log.WithContext(ctx).Infof("Add item to user[%v] cart: %v", userId, i)

	cart, err := uc.repo.GetUserCart(ctx, userId)
	if err != nil {
		return err
	}

	if cart == nil {
		cart = &Cart{
			UserId: userId,
			Items:  []Item{i},
		}
	} else {
		cart.Items = append(cart.Items, i)
	}

	err = uc.repo.UpdateUserCart(ctx, userId, cart)
	if err != nil {
		return err
	}

	return nil
}
