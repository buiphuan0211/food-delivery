package ginuser

import (
	"food-delivery/common"
	"food-delivery/component/appcontext"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Profile(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		u := c.MustGet(common.CurrentUser).(common.Requester)

		u.GetUserId()
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(u))
	}
}
