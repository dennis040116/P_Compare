package repository

import (
	"context"

	"github.com/amitshekhariitbhu/go-backend-clean-architecture/domain"
	"github.com/amitshekhariitbhu/go-backend-clean-architecture/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type cartRepository struct {
	database   mongo.Database
	collection string
}

func NewCartRepository(db mongo.Database, collection string) domain.CartRepository {
	return &cartRepository{
		database:   db,
		collection: collection,
	}
}

func (cr *cartRepository) CreateCart(c context.Context, userID primitive.ObjectID, product domain.Product) error {
	collection := cr.database.Collection(cr.collection)
	_, err := collection.InsertOne(c, bson.M{"user_id": userID, "product_id": product.ID})
	return err
}

func (cr *cartRepository) GetCart(c context.Context, userID primitive.ObjectID) ([]domain.ComparedCart, error) {
	collection := cr.database.Collection(cr.collection)
	var ccs []domain.ComparedCart
	cursor, err := collection.Find(c, bson.M{"user_id": userID})
	if err != nil {
		return ccs, err
	}
	err = cursor.All(c, &ccs)
	return ccs, err
}

func (cr *cartRepository) DeleteCartByProductID(c context.Context, userID primitive.ObjectID, productID primitive.ObjectID) error {
	collection := cr.database.Collection(cr.collection)
	_, err := collection.DeleteOne(c, bson.M{"user_id": userID, "product_id": productID})
	return err
}
