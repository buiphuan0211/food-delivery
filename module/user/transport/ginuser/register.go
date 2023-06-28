package ginuser

import (
	"food-delivery/common"
	"food-delivery/component/appcontext"
	"food-delivery/component/hasher"
	userbusiness "food-delivery/module/user/business"
	usermodel "food-delivery/module/user/model"
	userstore "food-delivery/module/user/store"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			ctx  = c.Request.Context()
			data usermodel.UserCreate
			db   = appCtx.GetMainDBConnection()
		)

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := userstore.NewSQLStore(db)

		md5 := hasher.NewMd5Hash()
		business := userbusiness.NewRegisterBusiness(store, md5)
		if err := business.Register(ctx, &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(gin.H{
			"id": data.ID,
		}))
	}
}
