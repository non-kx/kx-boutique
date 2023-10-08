package schema

import (
	"kx-boutique/app/shop/internal/biz"
	"time"
)

// Product schema.
type Product struct {
	ID          string    `bson:"_id"`
	CreatedAt   time.Time `bson:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at"`
	Name        string    `bson:"name"`
	Description string    `bson:"description"`
	ImageUrl    string    `bson:"image_url"`
	PriceTHB    int64     `bson:"price_thb"`
}

func (p *Product) ToBizProduct() *biz.Product {
	return &biz.Product{
		ID:          p.ID,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
		Name:        p.Name,
		Description: p.Description,
		ImageUrl:    p.ImageUrl,
		PriceTHB:    p.PriceTHB,
	}
}
