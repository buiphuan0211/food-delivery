package ginrstlike

import (
	"food-delivery/common"
	"food-delivery/component/appcontext"
	rstlikebiz "food-delivery/module/restaurantlike/biz"
	restaurantlikemodel "food-delivery/module/restaurantlike/model"
	restaurantlikestorage "food-delivery/module/restaurantlike/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// ListUsers return list user who liked the restaurant
func ListUsers(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		filter := restaurantlikemodel.Filter{
			RestaurantID: id,
		}

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fullfill()

		store := restaurantlikestorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := rstlikebiz.NewListUserLikeRestaurantBiz(store)

		result, err := biz.ListUsers(c.Request.Context(), &filter, &paging)
		if err != nil {
			return
		}

		//for i := range result {
		//	result[i].Mask(false)
		//}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))

	}
}
