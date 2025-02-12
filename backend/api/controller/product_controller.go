package controller

import (
	"log"
	"net/http"

	"github.com/amitshekhariitbhu/go-backend-clean-architecture/domain"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductController struct {
	ProductUsecase domain.ProductUsecase
}

func (pc *ProductController) Fetch(c *gin.Context) {
	userID, err := primitive.ObjectIDFromHex(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	// ob, err := primitive.ObjectIDFromHex("674986b1bf8f9ab72bae021b")
	// ud_products := []domain.Product{
	// 	{
	// 		ID:          ob,
	// 		PlatformID:  primitive.NewObjectID(),
	// 		Name:        "test",
	// 		Category:    "test",
	// 		Brand:       "test",
	// 		Description: "4466 the best product in 2025",
	// 		ImageURL:    "testt",
	// 		Price:       156,
	// 		Currency:    "100",
	// 		CreatedAt:   time.Now(),
	// 		UpdatedAt:   time.Now(),
	// 	},
	// }
	// err = pc.ProductUsecase.UpdateProducts(context.Background(), ud_products)

	log.Println("update")
	products, err := pc.ProductUsecase.Fetch(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (pc *ProductController) FetchByPlatform(c *gin.Context) {
	platformID, err := primitive.ObjectIDFromHex(c.Param("platform_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	userID, err := primitive.ObjectIDFromHex(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	products, err := pc.ProductUsecase.FetchByPlatform(c, userID, platformID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (pc *ProductController) Like(c *gin.Context) {
	productID, err := primitive.ObjectIDFromHex(c.Param("product_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	userID, err := primitive.ObjectIDFromHex(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	err = pc.ProductUsecase.LikeProduct(c, userID, productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, nil)
}

func (pc *ProductController) Dislike(c *gin.Context) {
	productID, err := primitive.ObjectIDFromHex(c.Param("product_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	userID, err := primitive.ObjectIDFromHex(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	err = pc.ProductUsecase.UnlikeProduct(c, userID, productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, nil)
}

func (pc *ProductController) FetchLikes(c *gin.Context) {
	userID, err := primitive.ObjectIDFromHex(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	likes, err := pc.ProductUsecase.FetchLikes(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, likes)
}

func (pc *ProductController) ShowDetail(c *gin.Context) {
	productID, err := primitive.ObjectIDFromHex(c.Param("product_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	product, price_his, err := pc.ProductUsecase.GetProductDetail(c, productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"product":       product,
		"price_history": price_his,
	})
}

func (pc *ProductController) Search(c *gin.Context) {

	userID, err := primitive.ObjectIDFromHex(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	query := c.Query("query")
	// if query == "" {
	// 	products, err := pc.ProductUsecase.Fetch(c, userID)
	// 	if err != nil {
	// 		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	// 		return
	// 	}
	// 	c.JSON(http.StatusOK, products)
	// }
	products, err := pc.ProductUsecase.SearchProducts(c, userID, query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}
