package main

import (
	"food-delivery/component/appcontext"
	"food-delivery/middleware"
	"food-delivery/module/restaurant/transport/ginrestaurant"
	"food-delivery/module/restaurantlike/transport/ginrstlike"
	"food-delivery/module/upload/transport/ginupload"
	"food-delivery/module/user/transport/ginuser"

	"github.com/gin-gonic/gin"
)

func setupRoute(appCtx appcontext.AppContext, v1 *gin.RouterGroup) {

	// Upload
	v1.POST("/upload", ginupload.Upload(appCtx))

	// Authenticate
	v1.POST("/register", ginuser.Register(appCtx))

	v1.POST("/login", ginuser.Login(appCtx))

	v1.GET("/profile", middleware.RequiredAuth(appCtx), ginuser.Profile(appCtx))

	// Restaurant
	rg := v1.Group("/restaurants", middleware.RequiredAuth(appCtx))

	rg.POST("", ginrestaurant.CreateRestaurant(appCtx))

	rg.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))

	rg.GET("", ginrestaurant.ListRestaurant(appCtx))

	rg.POST("/:id/liked-users", ginrstlike.UserLikeRestaurant(appCtx))

	rg.DELETE("/:id/liked-users", ginrstlike.UserDislikeRestaurant(appCtx))

	rg.GET("/:id/liked-users", ginrstlike.ListUsers(appCtx))

}
