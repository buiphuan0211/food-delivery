package main

import (
	"food-delivery/component/appcontext"
	"food-delivery/middleware"
	"food-delivery/module/user/transport/ginuser"

	"github.com/gin-gonic/gin"
)

func setupAdminRoute(appCtx appcontext.AppContext, v1 *gin.RouterGroup) {

	// Check roles
	admin := v1.Group("/admin",
		middleware.RequiredAuth(appCtx),
		middleware.RoleRequired(appCtx, "admin", "user"),
	)

	{
		admin.GET("/profile", ginuser.Profile(appCtx))
	}
}
