package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/amitshekhariitbhu/go-backend-clean-architecture/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type cartUsecase struct {
	cartRepository                domain.CartRepository
	productRepository             domain.ProductRepository
	productPriceHistoryRepository domain.ProductPriceHistoryRepository
	contextTimeout                time.Duration
}

func NewCartUsecase(cartRepository domain.CartRepository, productRepository domain.ProductRepository, productPriceHistoryRepository domain.ProductPriceHistoryRepository, timeout time.Duration) domain.CartUsecase {
	return &cartUsecase{
		cartRepository:                cartRepository,
		productRepository:             productRepository,
		productPriceHistoryRepository: productPriceHistoryRepository,
		contextTimeout:                timeout,
	}
}

func (cu *cartUsecase) Compare(c context.Context, productID primitive.ObjectID) (domain.ProductPriceHistory, error) {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()

	priceHistory, err := cu.productPriceHistoryRepository.GetProductPriceHistory(ctx, productID)
	if err != nil {
		return domain.ProductPriceHistory{}, fmt.Errorf("error fetching price history: %v", err)
	}
	return priceHistory, nil
}

func (cu *cartUsecase) Fetch(c context.Context, userID primitive.ObjectID) ([]domain.ComparedCartResponse, error) {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	productsInCart, err := cu.cartRepository.GetCart(ctx, userID)
	if err != nil {
		return nil, err
	}
	var ccrs []domain.ComparedCartResponse
	for _, product := range productsInCart {
		var ccr domain.ComparedCartResponse

		product, _ := cu.productRepository.GetProductByID(ctx, product.ProductID)
		ccr = domain.ComparedCartResponse{
			ProductID:   product.ID,
			PlatformID:  product.PlatformID,
			Description: product.Description,
			ImageURL:    product.ImageURL,
			Price:       product.Price,
		}
		ccrs = append(ccrs, ccr)
	}
	return ccrs, nil
}

func (cu *cartUsecase) Store(c context.Context, userID primitive.ObjectID, productID primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	product, err := cu.productRepository.GetProductByID(ctx, productID)
	if err != nil {
		return err
	}

	return cu.cartRepository.CreateCart(ctx, userID, product)
}

func (cu *cartUsecase) Delete(c context.Context, userID primitive.ObjectID, productID primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.cartRepository.DeleteCartByProductID(ctx, userID, productID)
}
