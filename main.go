package main

import (
	"food-delivery/component/appcontext"
	"food-delivery/middleware"
	"food-delivery/module/restaurant/transport/ginrestaurant"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func main() {
	dsn := "food_delivery:1234@tcp(127.0.0.1:3306)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	db = db.Debug()

	r := gin.Default()

	var appCtx = appcontext.NewAppContext(db)

	r.Use(middleware.Recover(appCtx))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Setup router ...
	v1 := r.Group("/v1")
	restaurantGroup := v1.Group("/restaurants")

	// Create
	restaurantGroup.POST("", ginrestaurant.CreateRestaurant(appCtx))

	// Delete
	restaurantGroup.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))

	// List
	restaurantGroup.GET("", ginrestaurant.ListRestaurant(appCtx))

	//// Detail
	//restaurantGroup.GET("/:id", func(c *gin.Context) {
	//	id, err := strconv.Atoi(c.Param("id"))
	//	if err != nil {
	//		c.JSON(http.StatusBadRequest, gin.H{
	//			"error": err,
	//		})
	//	}
	//
	//	var data Restaurant
	//	db.Where("id = ?", id).First(&data)
	//
	//	c.JSON(http.StatusOK, gin.H{
	//		"data": data,
	//	})
	//})
	//

	//	// Set default pagination
	//	if pageData.Page == 0 {
	//		pageData.Page = 1
	//	}
	//	if pageData.Limit == 0 {
	//		pageData.Limit = 5
	//	}
	//
	//	offset := (pageData.Page - 1) * pageData.Limit
	//
	//	db.Offset(offset).Order("id desc").Limit(pageData.Limit).Find(&data)
	//	c.JSON(http.StatusOK, gin.H{
	//		"list": data,
	//	})
	//})
	//
	//// Update ...
	//restaurantGroup.PATCH("/:id", func(c *gin.Context) {
	//	id, err := strconv.Atoi(c.Param("id"))
	//	if err != nil {
	//		c.JSON(http.StatusBadRequest, gin.H{
	//			"error": err.Error(),
	//		})
	//		return
	//	}
	//
	//	var payload RestaurantUpdate
	//	if err := c.ShouldBind(&payload); err != nil {
	//		c.JSON(http.StatusBadRequest, gin.H{
	//			"error": err.Error(),
	//		})
	//		return
	//	}
	//
	//	if err := db.Where("id = ?", id).Updates(&payload); err != nil {
	//		c.JSON(http.StatusBadRequest, gin.H{
	//			"error": err,
	//		})
	//		return
	//	}
	//	c.JSON(http.StatusOK, gin.H{
	//		"data": "success",
	//	})
	//})
	//

	r.Run()
}
