package biz

import (
	"context"
	"time"

	v1 "kx-boutique/api/shop/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrProductNotFound is product not found.
	ErrProductNotFound = errors.NotFound(v1.ErrorReason_PRODUCT_NOT_FOUND.String(), "product not found")
)

// Product model.
type Product struct {
	ID          string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	Description string
	ImageUrl    string
	PriceTHB    int64
}

// Products Paginate model
type ProductsPaginate struct {
	Products   []Product
	PageCount  int64
	TotalCount int64
}

// ProductRepo is a Product repo.
type ProductRepo interface {
	Save(context.Context, *Product) (*Product, error)
	Update(context.Context, *Product) (*Product, error)
	FindByID(ctx context.Context, id string) (*Product, error)
	ListAll(context.Context) ([]*Product, error)
	ListPaginate(ctx context.Context, page int64, limit int64) (*ProductsPaginate, error)
}

// ProductUsecase is a Product usecase.
type ProductUsecase struct {
	repo ProductRepo
	log  *log.Helper
}

// NewProductUsecase new a Product usecase.
func NewProductUsecase(repo ProductRepo, logger log.Logger) *ProductUsecase {
	return &ProductUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateProduct creates a Product, and returns the new Product.
func (uc *ProductUsecase) CreateProduct(ctx context.Context, p *Product) (*Product, error) {
	uc.log.WithContext(ctx).Infof("CreateProduct: %v", p)
	return uc.repo.Save(ctx, p)
}

// CreateProduct creates a Product, and returns the new Product.
func (uc *ProductUsecase) GetProductById(ctx context.Context, id string) (*Product, error) {
	uc.log.WithContext(ctx).Infof("Get product id: %v", id)

	product, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, ErrProductNotFound
	}

	uc.log.WithContext(ctx).Infof("Return product: %v", product)

	return product, nil
}

// CreateProduct creates a Product, and returns the new Product.
func (uc *ProductUsecase) GetProductsPaginate(ctx context.Context, page int64, limit int64) (*ProductsPaginate, error) {
	uc.log.WithContext(ctx).Infof("Get products paginate (page, limit): %v, %v", page, limit)

	pgnate, err := uc.repo.ListPaginate(ctx, page, limit)
	if err != nil {
		return nil, err
	}

	return pgnate, nil
}
