package product

type ProductFilter struct {
	Name        string   `json:"name"`
	ProductType string   `json:"product_type"`
	MinPrice    *float64 `json:"min_price"`
	MaxPrice    *float64 `json:"max_price"`
	MinRating   *float64 `json:"min_rating"`
}
