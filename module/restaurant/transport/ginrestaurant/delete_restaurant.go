package ginrestaurant

import (
	"food-delivery/common"
	"food-delivery/component/appcontext"
	restaurantbusiness "food-delivery/module/restaurant/business"
	restaurantstorage "food-delivery/module/restaurant/storgage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func DeleteRestaurant(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			ctx       = c.Request.Context()
			db        = appCtx.GetMainDBConnection()
			requester = c.MustGet(common.CurrentUser).(common.Requester)
		)

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantstorage.NewSQLStore(db)

		business := restaurantbusiness.NewDeleteRestaurantBusiness(store, requester)

		if err := business.Delete(ctx, id); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
