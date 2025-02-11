package main

import (
	"log"
	"os"
	"product_service/product"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	_ = godotenv.Load()

	dsn := "host=" + os.Getenv("DB_HOST") +
	" port=" + os.Getenv("DB_PORT") +
	" user=" + os.Getenv("DB_USER") +
	" password=" + os.Getenv("DB_PASSWORD") +
	" dbname=" + os.Getenv("DB_NAME") +
	" sslmode=disable"

	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil{
		log.Fatalf("Failed to connect db: %v", err)
	}

	if err := db.AutoMigrate(&product.Product{}); err != nil{
		log.Fatalf("Failed automigrate: %v", err)
	}

	productRepo := &product.ProductRepo{DB: db}
	productService := &product.ProductService{Repo: productRepo}
	productHandler := &product.ProductHandler{Service: productService}

	r := gin.Default()

	r.GET("/products", productHandler.GetAllProducts)
	r.POST("/products/filter", productHandler.FilterProduct)

	adminRoutes := r.Group("/admin")
	adminRoutes.Use(product.AdminMiddleware())
	adminRoutes.POST("/products", productHandler.CreateProduct)

	port := os.Getenv("SERVER_PORT")
	if err := r.Run(":" + port); err != nil{
		log.Fatalf("Cannot to run the server: %v", err)
	}
}