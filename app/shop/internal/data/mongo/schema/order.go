package schema

import "time"

// Order schema.
type Order struct {
	ID        string    `bson:"_id"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
	ProductID string    `bson:"product_id"`
}
