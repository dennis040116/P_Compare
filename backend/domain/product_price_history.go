package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionProductPriceHistory = "product_price_history"
)

type ProductPriceHistory struct {
	ID           primitive.ObjectID     `bson:"_id,omitempty"` // MongoDB ID
	ProductID    primitive.ObjectID     `bson:"product_id"`    // 引用 Product
	PriceHistory map[string]interface{} `bson:"price_history"` // 价格历史记录
}

type ProductPriceHistoryRepository interface {
	GetProductPriceHistory(c context.Context, productID primitive.ObjectID) (ProductPriceHistory, error)
	CreateProductPriceHistory(c context.Context, productID primitive.ObjectID, priceHistory map[string]interface{}) error
	UpdateProductPriceHistory(c context.Context, productID primitive.ObjectID, timestamp string, price float64) error
}
