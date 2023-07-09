package ginrstlike

import (
	"food-delivery/common"
	"food-delivery/component/appcontext"
	restaurantstorage "food-delivery/module/restaurant/storgage"
	rstlikebiz "food-delivery/module/restaurantlike/biz"
	restaurantlikestorage "food-delivery/module/restaurantlike/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// UserDislikeRestaurant DELETE /v1/restaurants/:id/dislike
func UserDislikeRestaurant(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		id, err := strconv.Atoi(idStr)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantlikestorage.NewSQLStore(appCtx.GetMainDBConnection())
		decStore := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := rstlikebiz.NewUserDislikeRestaurantBiz(store, decStore)

		if err := biz.UserDislikeRestaurant(c.Request.Context(), requester.GetUserId(), id); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
