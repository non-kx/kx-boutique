package data

import (
	"context"
	"math"

	"kx-boutique/app/shop/internal/biz"
	"kx-boutique/app/shop/internal/data/mongo/schema"

	"github.com/go-kratos/kratos/v2/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const PRODUCT_COLLECTION_NAME = "products"

type productRepo struct {
	col *mongo.Collection
	log *log.Helper
}

// NewProductRepo .
func NewProductRepo(data *Data, logger log.Logger) biz.ProductRepo {
	return &productRepo{
		col: data.db.Collection(PRODUCT_COLLECTION_NAME),
		log: log.NewHelper(logger),
	}
}

func (r *productRepo) Save(ctx context.Context, p *biz.Product) (*biz.Product, error) {
	return p, nil
}

func (r *productRepo) Update(ctx context.Context, p *biz.Product) (*biz.Product, error) {
	return p, nil
}

func (r *productRepo) FindByID(ctx context.Context, id string) (*biz.Product, error) {
	var product schema.Product

	docId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.D{{Key: "_id", Value: docId}}
	result := r.col.FindOne(ctx, filter)

	r.log.Infof("Raw from mongo: %v", result)

	err = result.Decode(&product)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		} else {
			return nil, err
		}
	}

	r.log.Debug(product)

	return &biz.Product{
		ID:          product.ID,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
		Name:        product.Name,
		Description: product.Description,
		ImageUrl:    product.ImageUrl,
	}, nil
}

func (r *productRepo) ListAll(context.Context) ([]*biz.Product, error) {
	return nil, nil
}

func (r *productRepo) ListPaginate(ctx context.Context, page int64, limit int64) (*biz.ProductsPaginate, error) {
	var products []schema.Product

	opts := options.Count().SetHint("_id_")
	count, err := r.col.CountDocuments(context.TODO(), bson.D{}, opts)
	if err != nil {
		return nil, err
	}

	pagecount := int64(math.Ceil(float64(count) / float64(limit)))

	skip := (page - 1) * limit
	skipopts := options.FindOptions{Limit: &limit, Skip: &skip}
	cursor, err := r.col.Find(ctx, bson.D{}, &skipopts)
	defer cursor.Close(ctx)

	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &products)
	if err != nil {
		return nil, err
	}

	bizproduct := make([]biz.Product, 0)
	for _, prod := range products {
		bizproduct = append(bizproduct, *prod.ToBizzProduct())
	}

	return &biz.ProductsPaginate{
		Products:   bizproduct,
		PageCount:  pagecount,
		TotalCount: count,
	}, nil
}
