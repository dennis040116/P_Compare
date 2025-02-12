package repository

import (
	"context"
	"fmt"

	"github.com/amitshekhariitbhu/go-backend-clean-architecture/domain"
	"github.com/amitshekhariitbhu/go-backend-clean-architecture/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type productPriceHistoryRepository struct {
	database   mongo.Database
	collection string
}

func NewProductPriceHistoryRepository(db mongo.Database, collection string) domain.ProductPriceHistoryRepository {
	return &productPriceHistoryRepository{
		database:   db,
		collection: collection,
	}
}

func (pphr *productPriceHistoryRepository) GetProductPriceHistory(c context.Context, productID primitive.ObjectID) (domain.ProductPriceHistory, error) {
	collection := pphr.database.Collection(pphr.collection)
	var pph domain.ProductPriceHistory
	err := collection.FindOne(c, bson.M{"product_id": productID}).Decode(&pph)
	return pph, err
}

func (pphr *productPriceHistoryRepository) CreateProductPriceHistory(c context.Context, productID primitive.ObjectID, priceHistory map[string]interface{}) error {
	collection := pphr.database.Collection(pphr.collection)
	_, err := collection.InsertOne(c, bson.M{"product_id": productID, "price_history": priceHistory})
	return err
}

func (pphr *productPriceHistoryRepository) UpdateProductPriceHistory(c context.Context, productID primitive.ObjectID, timestamp string, price float64) error {
	collection := pphr.database.Collection(pphr.collection)
	filter := bson.M{"product_id": productID}
	update := bson.M{
		"$set": bson.M{
			fmt.Sprintf("price_history.%s", timestamp): price, // 使用动态的字段名（时间戳）
		},
	}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}
