package ginrstlike

import (
	"food-delivery/common"
	"food-delivery/component/appcontext"
	restaurantstorage "food-delivery/module/restaurant/storgage"
	rstlikebiz "food-delivery/module/restaurantlike/biz"
	restaurantlikemodel "food-delivery/module/restaurantlike/model"
	restaurantlikestorage "food-delivery/module/restaurantlike/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// UserLikeRestaurant /v1/restaurants/:id/dislike
func UserLikeRestaurant(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		id, err := strconv.Atoi(idStr)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		data := restaurantlikemodel.Like{
			RestaurantID: id,
			UserID:       requester.GetUserId(),
		}

		store := restaurantlikestorage.NewSQLStore(appCtx.GetMainDBConnection())
		incStore := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := rstlikebiz.NewUserLikeRestaurantBiz(store, incStore)

		if err := biz.LikeRestaurant(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
