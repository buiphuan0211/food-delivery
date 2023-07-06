package main

import (
	"food-delivery/component/appcontext"
	"food-delivery/component/uploadprovider"
	"food-delivery/middleware"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

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

	r.Use(middleware.Recover(appCtx))

	// Nếu truy cập /static thì gin sẽ đi kiếm thư mục "./static" đọc vô
	r.Static("/static", "./static")

	// Setup router ...
	v1 := r.Group("/v1")

	v1.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	setupRoute(appCtx, v1)
	setupAdminRoute(appCtx, v1)

	r.Run()
}
