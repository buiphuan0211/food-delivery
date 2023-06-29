package middleware

import (
	"errors"
	"food-delivery/common"
	"food-delivery/component/appcontext"

	"github.com/gin-gonic/gin"
)

func RoleRequired(appCtx appcontext.AppContext, allowRoles ...string) func(c *gin.Context) {
	return func(c *gin.Context) {
		u := c.MustGet(common.CurrentUser).(common.Requester)

		for _, item := range allowRoles {
			if u.GetRole() == item {
				c.Next()
				// break
				return
			}
		}

		panic(common.ErrNoPermission(errors.New("invalid role user")))
	}
}
