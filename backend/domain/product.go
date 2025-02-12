package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionProduct = "products"
)

type Product struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"` // MongoDB ID
	PlatformID  primitive.ObjectID `bson:"platform_id"`   // 引用 EcommercePlatform
	Name        string             `bson:"name"`          // 商品名称
	Category    string             `bson:"category"`      // 商品类别
	Brand       string             `bson:"brand"`         // 品牌
	Description string             `bson:"description"`   // 商品描述
	ImageURL    string             `bson:"image_url"`     // 商品图片链接
	Price       float64            `bson:"price"`         // 商品价格
	Currency    string             `bson:"currency"`      // 货币单位
	CreatedAt   time.Time          `bson:"created_at"`    // 创建时间
	UpdatedAt   time.Time          `bson:"updated_at"`    // 更新时间
}

type ProductResponse struct {
	ProductID   primitive.ObjectID `json:"product_id"`
	PlatformID  primitive.ObjectID `bson:"platform_id"`
	Description string             `bson:"description"`  // 商品描述
	ImageURL    string             `bson:"image_url"`    // 商品图片链接
	Price       float64            `bson:"price"`        // 商品价格
	IsFavorited bool               `bson:"is_favorited"` // 是否被收藏
}

type ProductRepository interface {
	CreateProduct(c context.Context, product Product) error
	GetProductByID(c context.Context, productID primitive.ObjectID) (Product, error)
	UpdateProduct(c context.Context, product Product) error
	DeleteProduct(c context.Context, productID primitive.ObjectID) error
	FetchProducts(c context.Context) ([]Product, error)
	FetchProductsByPlatform(c context.Context, platformID primitive.ObjectID) ([]Product, error)

	SearchProducts(c context.Context, keyword string) ([]Product, error)
}

type ProductUsecase interface {
	CreateProducts(c context.Context, products []Product) error
	UpdateProducts(c context.Context, products []Product) error
	Fetch(c context.Context, userID primitive.ObjectID) ([]ProductResponse, error)
	FetchByPlatform(c context.Context, userID primitive.ObjectID, platformID primitive.ObjectID) ([]ProductResponse, error)
	FetchLikes(c context.Context, userID primitive.ObjectID) ([]ProductResponse, error)
	SearchProducts(c context.Context, userID primitive.ObjectID, keyword string) ([]ProductResponse, error)
	LikeProduct(c context.Context, userID primitive.ObjectID, productID primitive.ObjectID) error
	UnlikeProduct(c context.Context, userID primitive.ObjectID, productID primitive.ObjectID) error
	CheckPriceDropInFavorites(c context.Context, userID primitive.ObjectID) error
	CheckPriceDrop(c context.Context, userID primitive.ObjectID, productPrice float64, threshold float64) error
	GetProductDetail(c context.Context, productID primitive.ObjectID) (Product, ProductPriceHistory, error)
}
