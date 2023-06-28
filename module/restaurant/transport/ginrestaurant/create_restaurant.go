package ginrestaurant

import (
	"food-delivery/common"
	"food-delivery/component/appcontext"
	restaurantbusiness "food-delivery/module/restaurant/business"
	restaurantmodel "food-delivery/module/restaurant/model"
	restaurantstorage "food-delivery/module/restaurant/storgage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateRestaurant(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			ctx  = c.Request.Context()
			data restaurantmodel.RestaurantCreate
			db   = appCtx.GetMainDBConnection()
		)

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		data.UserId = requester.GetUserId()

		store := restaurantstorage.NewSQLStore(db)

		business := restaurantbusiness.NewCreateRestaurantBusiness(store)

		if err := business.CreateRestaurant(ctx, &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(gin.H{
			"id": data.ID,
		}))
	}
}
