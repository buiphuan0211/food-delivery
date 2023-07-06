package ginrestaurant

import (
	"food-delivery/common"
	"food-delivery/component/appcontext"
	restaurantbiz "food-delivery/module/restaurant/business"
	restaurantmodel "food-delivery/module/restaurant/model"
	restaurantrepo "food-delivery/module/restaurant/repository"
	restaurantstorage "food-delivery/module/restaurant/storgage"
	restaurantlikestorage "food-delivery/module/restaurantlike/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListRestaurant(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			ctx = c.Request.Context()
			db  = appCtx.GetMainDBConnection()
		)

		var pagingData common.Paging
		if err := c.ShouldBind(&pagingData); err != nil {
			panic(err)
		}

		pagingData.Fullfill()

		var filter restaurantmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(err)
		}

		store := restaurantstorage.NewSQLStore(db)
		likeStore := restaurantlikestorage.NewSQLStore(db)
		repo := restaurantrepo.NewListRestaurantRepo(store, likeStore)
		business := restaurantbiz.NewListRestaurantBusiness(repo)

		result, err := business.ListRestaurant(ctx, &filter, &pagingData)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, pagingData, filter))
	}
}
