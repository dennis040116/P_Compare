package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionFavoriteProduct = "favourite_products"
)

type FavoriteProduct struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"` // MongoDB ID
	UserID    primitive.ObjectID `bson:"user_id"`       // 引用 User
	ProductID primitive.ObjectID `bson:"product_id"`    // 引用 Product
	Threshold float64            `bson:"threshold"`
	CreatedAt time.Time          `bson:"created_at"` // 创建时间
}

type FavoriteProductRepository interface {
	CreateFavoriteProduct(c context.Context, userID primitive.ObjectID, productID primitive.ObjectID, threshold float64) error
	DeleteFavoriteProduct(c context.Context, userID primitive.ObjectID, productID primitive.ObjectID) error
	GetFavoriteProducts(c context.Context, userID primitive.ObjectID) ([]FavoriteProduct, error)
	IsFavorite(c context.Context, userID primitive.ObjectID, productID primitive.ObjectID) bool
}
