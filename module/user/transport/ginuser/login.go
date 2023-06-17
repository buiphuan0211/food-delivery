package ginuser

import (
	"food-delivery/common"
	"food-delivery/component/appcontext"
	"food-delivery/component/hasher"
	"food-delivery/component/tokenprovider/jwt"
	userbusiness "food-delivery/module/user/business"
	usermodel "food-delivery/module/user/model"
	userstore "food-delivery/module/user/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			ctx           = c.Request.Context()
			loginUserData usermodel.UserLogin
			db            = appCtx.GetMainDBConnection()
		)

		if err := c.ShouldBind(&loginUserData); err != nil {
			panic(err)
		}

		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())

		store := userstore.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()

		business := userbusiness.NewLoginBusiness(store, tokenProvider, md5, 60*60*24*30)

		token, err := business.Login(ctx, loginUserData)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(gin.H{
			"token": token.Token,
		}))
	}
}
