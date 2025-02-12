package usecase

import (
	"context"
	"fmt"
	"log"
	"net/smtp"
	"time"

	"github.com/amitshekhariitbhu/go-backend-clean-architecture/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type productUsecase struct {
	productRepository             domain.ProductRepository
	favoriteProductRepository     domain.FavoriteProductRepository
	productPriceHistoryRepository domain.ProductPriceHistoryRepository
	userRepository                domain.UserRepository
	contextTimeout                time.Duration
}

func NewProductUsecase(productRepository domain.ProductRepository, favoriteProductRepository domain.FavoriteProductRepository, productPriceHistoryRepository domain.ProductPriceHistoryRepository, userRepository domain.UserRepository, timeout time.Duration) domain.ProductUsecase {
	return &productUsecase{
		productRepository:             productRepository,
		favoriteProductRepository:     favoriteProductRepository,
		productPriceHistoryRepository: productPriceHistoryRepository,
		userRepository:                userRepository,
		contextTimeout:                timeout,
	}
}

func (pu *productUsecase) SearchProducts(c context.Context, userID primitive.ObjectID, keyword string) ([]domain.ProductResponse, error) {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	products, err := pu.productRepository.SearchProducts(ctx, keyword)
	if err != nil {
		return nil, fmt.Errorf("error fetching products: %v", err)
	}
	var prs []domain.ProductResponse
	for _, product := range products {
		pr := domain.ProductResponse{
			ProductID:   product.ID,
			PlatformID:  product.PlatformID,
			Description: product.Description,
			ImageURL:    product.ImageURL,
			Price:       product.Price,
			IsFavorited: pu.favoriteProductRepository.IsFavorite(ctx, userID, product.ID),
		}
		prs = append(prs, pr)
	}
	return prs, nil
}

func SendEmail(to string, subject string, body string) error {
	smtphost := "smtp.zju.edu.cn"
	smtpport := "25"
	sender := "3220300115@zju.edu.cn"
	password := "badboy2468888"
	log.Println("Sending email to: " + to)

	msg := "From: " + sender + "\n" + "To: " + to + "\n" + "Subject: " + subject + "\n" + body

	auth := smtp.PlainAuth("", sender, password, smtphost)

	err := smtp.SendMail(smtphost+":"+smtpport, auth, sender, []string{to}, []byte(msg))
	if err != nil {
		return fmt.Errorf("error sending email in SendEmail: %v", err)
	}
	return nil
}

func (pu *productUsecase) CheckPriceDrop(c context.Context, userID primitive.ObjectID, productPrice float64, threshold float64) error {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	if productPrice < threshold {

		userID := userID.Hex()
		user, err := pu.userRepository.GetByID(ctx, userID)
		log.Println("user: ", user)
		if err != nil {
			// log.Println("in err")
			return fmt.Errorf("error fetching user: %v", err)
		}
		to := user.Email

		// to := "3220300115@zju.edu.cn"
		subject := "Price Drop Alert"
		body := "The price of the product you are following has dropped below the threshold you set. Please check the product page for more details."
		err = SendEmail(to, subject, body)

		if err != nil {
			return fmt.Errorf("error sending email in CheckPriceDrop: %v", err)
		}
	}
	return nil
}

func (pu *productUsecase) CheckPriceDropInFavorites(c context.Context, userID primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	var fps []domain.FavoriteProduct
	fps, err := pu.favoriteProductRepository.GetFavoriteProducts(ctx, userID)
	if err != nil {
		return fmt.Errorf("error fetching favorite products: %v", err)
	}
	for _, fp := range fps {
		var product domain.Product
		product, err = pu.productRepository.GetProductByID(ctx, fp.ProductID)
		if err != nil {
			return fmt.Errorf("error fetching product: %v", err)
		}
		pu.CheckPriceDrop(ctx, userID, product.Price, fp.Threshold)
	}
	return nil
}

func (pu *productUsecase) UpdateProducts(c context.Context, products []domain.Product) error {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	for _, product := range products {
		err := pu.productRepository.UpdateProduct(ctx, product)
		if err != nil {
			return err
		}

		currentTime := time.Now()
		timeString := currentTime.Format("2006-01-02 15:04:05")

		priceHistory := make(map[string]interface{})
		priceHistory[timeString] = product.Price
		err = pu.productPriceHistoryRepository.UpdateProductPriceHistory(ctx, product.ID, timeString, product.Price)
		if err != nil {
			return err
		}
	}

	return nil
}

func (pu *productUsecase) CreateProducts(c context.Context, products []domain.Product) error {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	log.Println("in usecase create products")
	defer cancel()
	for _, product := range products {
		err := pu.productRepository.CreateProduct(ctx, product)
		if err != nil {
			return err
		}

		currentTime := time.Now()
		timeString := currentTime.Format("2006-01-02 15:04:05")

		priceHistory := make(map[string]interface{})
		priceHistory[timeString] = product.Price
		err = pu.productPriceHistoryRepository.CreateProductPriceHistory(ctx, product.ID, priceHistory)
		if err != nil {
			return err
		}
	}

	return nil
}

func (pu *productUsecase) DeleteProduct(c context.Context, productID primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	err := pu.productRepository.DeleteProduct(ctx, productID)
	if err != nil {
		return err
	}
	// err = pu.productPriceHistoryRepository.DeleteProductPriceHistory(ctx, productID)
	// if err != nil {
	// 	return err
	// }
	// err = pu.favoriteProductRepository.DeleteFavoriteProduct(ctx, userID,productID)
	// if err != nil {
	// 	return err
	// }
	return nil
}

func (pu *productUsecase) FetchByPlatform(c context.Context, userID primitive.ObjectID, platformID primitive.ObjectID) ([]domain.ProductResponse, error) {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	products, err := pu.productRepository.FetchProductsByPlatform(ctx, platformID)
	if err != nil {
		return nil, fmt.Errorf("error fetching products: %v", err)
	}
	var prs []domain.ProductResponse
	for _, product := range products {
		pr := domain.ProductResponse{
			ProductID:   product.ID,
			PlatformID:  product.PlatformID,
			Description: product.Description,
			ImageURL:    product.ImageURL,
			Price:       product.Price,
			IsFavorited: pu.favoriteProductRepository.IsFavorite(ctx, userID, product.ID),
		}
		prs = append(prs, pr)
	}
	return prs, nil
}

func (pu *productUsecase) Fetch(c context.Context, userID primitive.ObjectID) ([]domain.ProductResponse, error) {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	products, err := pu.productRepository.FetchProducts(ctx)
	if err != nil {
		return nil, fmt.Errorf("error fetching products: %v", err)
	}
	var prs []domain.ProductResponse
	for _, product := range products {

		pr := domain.ProductResponse{
			ProductID:   product.ID,
			PlatformID:  product.PlatformID,
			Description: product.Description,
			ImageURL:    product.ImageURL,
			Price:       product.Price,
			IsFavorited: pu.favoriteProductRepository.IsFavorite(ctx, userID, product.ID),
		}
		prs = append(prs, pr)
	}
	return prs, nil
}

func (pu *productUsecase) FetchLikes(c context.Context, userID primitive.ObjectID) ([]domain.ProductResponse, error) {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	var fps []domain.FavoriteProduct
	fps, err := pu.favoriteProductRepository.GetFavoriteProducts(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("error fetching favorite products: %v", err)
	}
	var prs []domain.ProductResponse
	for _, fp := range fps {
		var product domain.Product
		product, err = pu.productRepository.GetProductByID(ctx, fp.ProductID)
		if err != nil {
			return nil, fmt.Errorf("error fetching product: %v", err)
		}
		pr := domain.ProductResponse{
			ProductID:   product.ID,
			PlatformID:  product.PlatformID,
			Description: product.Description,
			ImageURL:    product.ImageURL,
			Price:       product.Price,
			IsFavorited: true,
		}
		prs = append(prs, pr)
	}
	return prs, nil
}

func (pu *productUsecase) GetProductDetail(c context.Context, productID primitive.ObjectID) (domain.Product, domain.ProductPriceHistory, error) {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	product, err := pu.productRepository.GetProductByID(ctx, productID)
	if err != nil {
		return domain.Product{}, domain.ProductPriceHistory{}, fmt.Errorf("error fetching product: %v", err)
	}
	priceHistory, err := pu.productPriceHistoryRepository.GetProductPriceHistory(ctx, productID)
	if err != nil {
		return domain.Product{}, domain.ProductPriceHistory{}, fmt.Errorf("error fetching price history: %v", err)
	}
	return product, priceHistory, nil
}

func (pu *productUsecase) LikeProduct(c context.Context, userID primitive.ObjectID, productID primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	product, err := pu.productRepository.GetProductByID(ctx, productID)
	if err != nil {
		return fmt.Errorf("error fetching product: %v", err)
	}
	return pu.favoriteProductRepository.CreateFavoriteProduct(ctx, userID, productID, product.Price)
}

func (pu *productUsecase) UnlikeProduct(c context.Context, userID primitive.ObjectID, productID primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	return pu.favoriteProductRepository.DeleteFavoriteProduct(ctx, userID, productID)
}

// func (pu *productUsecase) GetPriceHistory(c context.Context, productID primitive.ObjectID) (domain.ProductPriceHistory, error) {
// 	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
// 	defer cancel()
// 	return pu.productRepository.GetProductPriceHistory(ctx, productID)
// }

// func (pu *productUsecase) GetProducts(c context.Context) ([]domain.Product, error) {
// 	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
// 	defer cancel()
// 	return pu.productRepository.FetchProducts(ctx)
// }

// func (pu *productUsecase) SearchProducts(c context.Context, keyword string) ([]domain.Product, error) {
// 	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
// 	defer cancel()

// 	return pu.productRepository.SearchProducts(ctx, keyword)
// }
