package ginuser

import (
	"food-delivery/common"
	"food-delivery/component/appcontext"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Profile(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		u := c.MustGet(common.CurrentUser).(common.Requester)

		u.GetUserId()
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(u))
	}
}
