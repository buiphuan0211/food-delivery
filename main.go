package main

import (
	"food-delivery/component/appcontext"
	"food-delivery/component/uploadprovider"
	"food-delivery/module/restaurant/transport/ginrestaurant"
	"food-delivery/module/upload/transport/ginupload"
	"food-delivery/module/user/transport/ginuser"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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

	// Nếu truy cập /static thì gin sẽ đi kiếm thư mục "./static" đọc vô
	r.Static("/static", "./static")

	// Load .env file
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	var (
		s3BucketName = os.Getenv("S3_BUCKET_NAME")
		s3Region     = os.Getenv("S3_REGION")
		s3APIkey     = os.Getenv("S3_API_KEY")
		s3SecretKey  = os.Getenv("S3_SECRET_KEY")
		s3Domain     = os.Getenv("S3_DOMAIN")
		secretKey    = os.Getenv("SECRET_KEY")
		s3Provider   = uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIkey, s3SecretKey, s3Domain)
		appCtx       = appcontext.NewAppContext(db, s3Provider, secretKey)
	)

	// Setup router ...
	v1 := r.Group("/v1")

	v1.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Upload
	v1.POST("/upload", ginupload.Upload(appCtx))

	// Authenticate
	v1.POST("/register", ginuser.Register(appCtx))

	v1.POST("/login", ginuser.Login(appCtx))

	// Restaurant
	restaurantGroup := v1.Group("/restaurants")

	restaurantGroup.POST("", ginrestaurant.CreateRestaurant(appCtx))

	restaurantGroup.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))

	restaurantGroup.GET("", ginrestaurant.ListRestaurant(appCtx))

	r.Run()
}
