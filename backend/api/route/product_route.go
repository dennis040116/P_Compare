package route

import (
	"time"

	"github.com/amitshekhariitbhu/go-backend-clean-architecture/api/controller"
	"github.com/amitshekhariitbhu/go-backend-clean-architecture/bootstrap"
	"github.com/amitshekhariitbhu/go-backend-clean-architecture/domain"
	"github.com/amitshekhariitbhu/go-backend-clean-architecture/mongo"
	"github.com/amitshekhariitbhu/go-backend-clean-architecture/repository"
	"github.com/amitshekhariitbhu/go-backend-clean-architecture/usecase"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gin-gonic/gin"
)

func NewProductRouter(env *bootstrap.Env, timeout time.Duration, es *elasticsearch.Client, db mongo.Database, group *gin.RouterGroup) {
	pr := repository.NewProductRepository(db, es, domain.CollectionProduct)
	fps := repository.NewFavoriteProductRepository(db, domain.CollectionFavoriteProduct)
	pph := repository.NewProductPriceHistoryRepository(db, domain.CollectionProductPriceHistory)
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	pc := &controller.ProductController{
		ProductUsecase: usecase.NewProductUsecase(pr, fps, pph, ur, timeout),
	}

	group.GET("/products/:user_id", pc.Fetch)
	group.GET("/products/:user_id/:platform_id", pc.FetchByPlatform)
	group.PUT("/products/like/:user_id/:product_id", pc.Like)
	group.PUT("/products/dislike/:user_id/:product_id", pc.Dislike)
	group.GET("/products/likes/:user_id", pc.FetchLikes)
	group.GET("/products/details/:user_id/:product_id", pc.ShowDetail)
	group.GET("/products/search/:user_id", pc.Search)
}
