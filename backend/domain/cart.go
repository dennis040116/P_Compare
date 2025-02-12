package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionCart = "carts"
)

type ComparedCart struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"` // MongoDB ID
	ProductID primitive.ObjectID `bson:"product_id"`    // 引用 Product
	UserID    primitive.ObjectID `bson:"user_id"`       // 引用 User
}

type ComparedCartResponse struct {
	ProductID   primitive.ObjectID `json:"product_id"`
	PlatformID  primitive.ObjectID `bson:"platform_id"` // 引用 EcommercePlatform
	Description string             `bson:"description"` // 商品描述
	ImageURL    string             `bson:"image_url"`   // 商品图片链接
	Price       float64            `bson:"price"`       // 商品价格
}

type CartRepository interface {
	CreateCart(c context.Context, userID primitive.ObjectID, product Product) error
	GetCart(c context.Context, userID primitive.ObjectID) ([]ComparedCart, error)
	DeleteCartByProductID(c context.Context, userID primitive.ObjectID, productID primitive.ObjectID) error
}

type CartUsecase interface {
	Compare(c context.Context, productID primitive.ObjectID) (ProductPriceHistory, error)
	Fetch(c context.Context, userID primitive.ObjectID) ([]ComparedCartResponse, error)
	Store(c context.Context, userID primitive.ObjectID, productID primitive.ObjectID) error
	Delete(c context.Context, userID primitive.ObjectID, productID primitive.ObjectID) error
}
