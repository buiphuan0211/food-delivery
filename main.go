package main

import (
	"food-delivery/component/appcontext"
	"food-delivery/component/uploadprovider"
	"food-delivery/middleware"
	"food-delivery/module/restaurant/transport/ginrestaurant"
	"food-delivery/module/upload/transport/ginupload"
	"food-delivery/module/user/transport/ginuser"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "food_delivery:1234@tcp(127.0.0.1:3306)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	db = db.Debug()

	r := gin.Default()

	s3bucketName := "food-delivery-img"
	s3Region := "ap-southeast-2"
	s3APIkey := "AKIA436TMGDHAGNJ6NNR"
	s3SecretKey := "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"
	s3Domain := "https://dm83ozfygdntq.cloudfront.net"
	secretKey := "MY_SECRET_KEY"

	s3Provider := uploadprovider.NewS3Provider(s3bucketName, s3Region, s3APIkey, s3SecretKey, s3Domain)

	var appCtx = appcontext.NewAppContext(db, s3Provider, secretKey)

	r.Use(middleware.Recover(appCtx))

	// Nếu truy cập /static thì gin sẽ đi kiếm thư mục "./static" đọc vô
	r.Static("/static", "./static")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Setup router ...
	v1 := r.Group("/v1")

	// Upload -------
	v1.POST("/upload", ginupload.Upload(appCtx))

	// Authenticate  -------
	v1.POST("/register", ginuser.Register(appCtx))

	v1.POST("/login", ginuser.Login(appCtx))

	// Restaurant -------
	restaurantGroup := v1.Group("/restaurants")

	restaurantGroup.POST("", ginrestaurant.CreateRestaurant(appCtx))

	restaurantGroup.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))

	restaurantGroup.GET("", ginrestaurant.ListRestaurant(appCtx))

	r.Run()
}
