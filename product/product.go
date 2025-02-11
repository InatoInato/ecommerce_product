package product

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name string `json:"name"`
	ProductType string `json:"product_type"`
	ImageUrl string `json:"image_url"`
	Price float64 `json:"price"`
	Rating float64 `json:"rating"`
	Discount float64 `json:"discount"`
}