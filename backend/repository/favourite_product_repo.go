package repository

import (
	"context"
	"time"

	"github.com/amitshekhariitbhu/go-backend-clean-architecture/domain"
	"github.com/amitshekhariitbhu/go-backend-clean-architecture/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type favoriteProductRepository struct {
	database   mongo.Database
	collection string
}

func NewFavoriteProductRepository(db mongo.Database, collection string) domain.FavoriteProductRepository {
	return &favoriteProductRepository{
		database:   db,
		collection: collection,
	}
}

func (fpr *favoriteProductRepository) CreateFavoriteProduct(c context.Context, userID primitive.ObjectID, productID primitive.ObjectID, threshold float64) error {
	collection := fpr.database.Collection(fpr.collection)
	_, err := collection.InsertOne(c, bson.M{"user_id": userID, "product_id": productID, "threshold": threshold, "created_at": time.Now()})
	return err
}

func (fpr *favoriteProductRepository) DeleteFavoriteProduct(c context.Context, userID primitive.ObjectID, productID primitive.ObjectID) error {
	collection := fpr.database.Collection(fpr.collection)
	_, err := collection.DeleteOne(c, bson.M{"user_id": userID, "product_id": productID})
	return err
}

func (fpr *favoriteProductRepository) GetFavoriteProducts(c context.Context, userID primitive.ObjectID) ([]domain.FavoriteProduct, error) {
	collection := fpr.database.Collection(fpr.collection)
	var fps []domain.FavoriteProduct
	cursor, err := collection.Find(c, bson.M{"user_id": userID})
	if err != nil {
		return fps, err
	}
	err = cursor.All(c, &fps)
	return fps, err
}

func (fpr *favoriteProductRepository) IsFavorite(c context.Context, userID primitive.ObjectID, productID primitive.ObjectID) bool {
	collection := fpr.database.Collection(fpr.collection)
	count, _ := collection.CountDocuments(c, bson.M{"user_id": userID, "product_id": productID})
	return count > 0
}
