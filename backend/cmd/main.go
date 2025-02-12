package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	route "github.com/amitshekhariitbhu/go-backend-clean-architecture/api/route"
	"github.com/amitshekhariitbhu/go-backend-clean-architecture/bootstrap"
	"github.com/amitshekhariitbhu/go-backend-clean-architecture/domain"
	"github.com/amitshekhariitbhu/go-backend-clean-architecture/mongo"
	"github.com/amitshekhariitbhu/go-backend-clean-architecture/repository"
	"github.com/amitshekhariitbhu/go-backend-clean-architecture/usecase"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {

	app := bootstrap.App()

	env := app.Env

	db := app.Mongo.Database(env.DBName)
	log.Println("Connected to Database(", env.DBName, ")!")

	es := app.ES

	defer app.CloseDBConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	createProduct(env, timeout, db, es)

	go processProductData(env, timeout, db, es)

	gin := gin.Default()

	gin.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // 允许的前端域名
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour, // 缓存时间
	}))

	route.Setup(env, timeout, db, es, gin)

	gin.Run(env.ServerAddress)
}

type C_Product struct {
	ID          string  `json:"product_id"`  // 引用 EcommercePlatform
	Name        string  `bson:"name"`        // 商品名称
	Category    string  `bson:"category"`    // 商品类别
	Brand       string  `bson:"brand"`       // 品牌
	Description string  `bson:"description"` // 商品描述
	ImageURL    string  `bson:"image_url"`   // 商品图片链接
	Price       float64 `bson:"price"`       // 商品价格
	Currency    string  `bson:"currency"`    // 货币单位
}

// 读取文件并解析数据
func readProductDataFromFile(filePath string) ([]domain.Product, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("unable to open file %s: %v", filePath, err)
	}
	defer file.Close()

	var products []C_Product
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&products)
	if err != nil {
		return nil, fmt.Errorf("unable to decode JSON from %s: %v", filePath, err)
	}

	var platform_id primitive.ObjectID
	if filePath == "./assets/dataset/taobao.json" {
		platform_id, _ = primitive.ObjectIDFromHex("675bc9d0960209b29601aea0")
	} else if filePath == "./assets/dataset/jingdong.json" {
		platform_id, _ = primitive.ObjectIDFromHex("675bc9d0960209b29601aea1")
	}

	var domainProducts []domain.Product
	for _, product := range products {
		log.Println(product)
		product_id, _ := primitive.ObjectIDFromHex(product.ID)
		domainProducts = append(domainProducts, domain.Product{
			ID:          product_id,
			PlatformID:  platform_id,
			Name:        product.Name,
			Category:    product.Category,
			Brand:       product.Brand,
			Description: product.Description,
			ImageURL:    product.ImageURL,
			Price:       product.Price,
			Currency:    product.Currency,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		})
	}

	return domainProducts, nil
}

// 定时执行的任务
func processProductData(env *bootstrap.Env, timeout time.Duration, db mongo.Database, es *elasticsearch.Client) {
	// 每小时执行一次任务
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	pr := repository.NewProductRepository(db, es, domain.CollectionProduct)
	fps := repository.NewFavoriteProductRepository(db, domain.CollectionFavoriteProduct)
	pph := repository.NewProductPriceHistoryRepository(db, domain.CollectionProductPriceHistory)
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	pu := usecase.NewProductUsecase(pr, fps, pph, ur, timeout)

	for {
		select {
		case <-ticker.C:
			// 读取 taobao.json 数据
			taobaoProducts, err := readProductDataFromFile("./assets/dataset/taobao.json")
			if err != nil {
				fmt.Println("Error reading Taobao products:", err)
			} else {
				// 处理 Taobao 数据
				fmt.Printf("Read %d products from Taobao\n", len(taobaoProducts))
				// 这里可以将数据存入数据库或进一步处理
				pu.UpdateProducts(context.Background(), taobaoProducts)
			}

			// 读取 京东.json 数据
			jdProducts, err := readProductDataFromFile("./assets/dataset/jingdong.json")
			if err != nil {
				fmt.Println("Error reading JD products:", err)
			} else {
				// 处理 京东 数据
				fmt.Printf("Read %d products from JD\n", len(jdProducts))
				// 这里可以将数据存入数据库或进一步处理
				pu.UpdateProducts(context.Background(), jdProducts)
			}
			users, _ := ur.Fetch(context.Background())

			for _, user := range users {
				pu.CheckPriceDropInFavorites(context.Background(), user.ID)
			}
		}
	}
}

func createProduct(env *bootstrap.Env, timeout time.Duration, db mongo.Database, es *elasticsearch.Client) {
	pr := repository.NewProductRepository(db, es, domain.CollectionProduct)
	fps := repository.NewFavoriteProductRepository(db, domain.CollectionFavoriteProduct)
	pph := repository.NewProductPriceHistoryRepository(db, domain.CollectionProductPriceHistory)
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	pu := usecase.NewProductUsecase(pr, fps, pph, ur, timeout)

	taobaoProducts, err := readProductDataFromFile("./assets/dataset/taobao.json")
	if err != nil {
		fmt.Println("Error reading Taobao products:", err)
	} else {
		// 处理 Taobao 数据
		fmt.Printf("Read %d products from Taobao\n", len(taobaoProducts))
		// 这里可以将数据存入数据库或进一步处理
		log.Println(taobaoProducts)
		pu.CreateProducts(context.Background(), taobaoProducts)
	}

	// 读取 京东.json 数据
	jdProducts, err := readProductDataFromFile("./assets/dataset/jingdong.json")
	if err != nil {
		fmt.Println("Error reading JD products:", err)
	} else {
		// 处理 京东 数据
		fmt.Printf("Read %d products from JD\n", len(jdProducts))
		// 这里可以将数据存入数据库或进一步处理
		pu.CreateProducts(context.Background(), jdProducts)
	}
}
