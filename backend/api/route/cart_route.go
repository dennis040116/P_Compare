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

func NewCartRouter(env *bootstrap.Env, timeout time.Duration, es *elasticsearch.Client, db mongo.Database, group *gin.RouterGroup) {
	cr := repository.NewCartRepository(db, domain.CollectionCart)
	pr := repository.NewProductRepository(db, es, domain.CollectionProduct)
	pph := repository.NewProductPriceHistoryRepository(db, domain.CollectionProductPriceHistory)
	cc := &controller.CartController{
		CartUsecase: usecase.NewCartUsecase(cr, pr, pph, timeout),
	}
	group.GET("/cart/:user_id", cc.Fetch)
	group.POST("/cart/:user_id/:product_id", cc.Store)
	group.GET("/cart/compare/:user_id", cc.Compare)
	group.DELETE("/cart/delete/:user_id/:product_id", cc.Delete)
}
