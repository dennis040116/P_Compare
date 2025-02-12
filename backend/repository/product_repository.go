package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/amitshekhariitbhu/go-backend-clean-architecture/domain"
	"github.com/amitshekhariitbhu/go-backend-clean-architecture/mongo"
	"github.com/elastic/go-elasticsearch/v8"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type productRepository struct {
	database   mongo.Database
	es         *elasticsearch.Client
	collection string
}

func NewProductRepository(db mongo.Database, es *elasticsearch.Client, collection string) domain.ProductRepository {
	return &productRepository{
		database:   db,
		es:         es,
		collection: collection,
	}
}

func (pr *productRepository) CreateProduct(c context.Context, product domain.Product) error {
	collection := pr.database.Collection(pr.collection)

	document := map[string]interface{}{
		"product_id":  product.ID,
		"product_des": product.Description,
	}
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(document); err != nil {
		log.Fatalf("Error encoding document: %s", err)
	}
	res, err := pr.es.Index(
		"product",
		&buf,
		pr.es.Index.WithContext(c),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	_, err = collection.InsertOne(c, product)
	return err
}

func (pr productRepository) FetchProducts(c context.Context) ([]domain.Product, error) {
	collection := pr.database.Collection(pr.collection)
	var products []domain.Product
	cursor, err := collection.Find(c, bson.M{})
	if err != nil {
		return products, err
	}
	err = cursor.All(c, &products)
	return products, err
}

func (pr *productRepository) FetchProductsByPlatform(c context.Context, platformID primitive.ObjectID) ([]domain.Product, error) {
	collection := pr.database.Collection(pr.collection)
	var products []domain.Product
	cursor, err := collection.Find(c, bson.M{"platform_id": platformID})
	if err != nil {
		return products, err
	}
	err = cursor.All(c, &products)
	return products, err
}

func (pr *productRepository) GetProductByID(c context.Context, productID primitive.ObjectID) (domain.Product, error) {
	collection := pr.database.Collection(pr.collection)
	var product domain.Product
	err := collection.FindOne(c, bson.M{"_id": productID}).Decode(&product)
	return product, err
}

func (pr *productRepository) UpdateProduct(c context.Context, product domain.Product) error {
	collection := pr.database.Collection(pr.collection)
	filter := bson.M{"_id": product.ID}
	update := bson.M{"$set": product}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func (pr *productRepository) DeleteProduct(c context.Context, productID primitive.ObjectID) error {
	collection := pr.database.Collection(pr.collection)

	// Delete from elasticsearch
	_, err := pr.es.Delete(
		"product",
		productID.Hex(),
		pr.es.Delete.WithContext(c),
	)
	if err != nil {
		log.Fatalf("Error deleting document: %s", err)
	}

	_, err = collection.DeleteOne(c, bson.M{"_id": productID})
	return err
}

func (pr *productRepository) SearchProducts(c context.Context, keyword string) ([]domain.Product, error) {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"product_des": keyword,
			},
		},
	}

	var buf bytes.Buffer

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, fmt.Errorf("error encoding query: %s", err)
	}

	res, err := pr.es.Search(
		pr.es.Search.WithContext(c),
		pr.es.Search.WithIndex("product"),
		pr.es.Search.WithBody(&buf),
		pr.es.Search.WithTrackTotalHits(true),
	)

	if err != nil {
		return nil, fmt.Errorf("error getting response: %s", err)
	}

	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		return nil, fmt.Errorf("[%s] %s: %s",
			res.Status(),
			e["error"].(map[string]interface{})["type"],
			e["error"].(map[string]interface{})["reason"])

	}

	var products []domain.Product
	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, fmt.Errorf("error parsing the response body: %s", err)
	}

	hits := r["hits"].(map[string]interface{})["hits"].([]interface{})

	for _, hit := range hits {
		source := hit.(map[string]interface{})["_source"]
		var product domain.Product
		product.ID, err = primitive.ObjectIDFromHex(source.(map[string]interface{})["product_id"].(string))
		if err != nil {
			return nil, err
		}
		product, _ = pr.GetProductByID(c, product.ID)
		products = append(products, product)
	}

	return products, nil
}
