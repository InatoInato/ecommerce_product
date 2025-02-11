package product

type ProductService struct {
	Repo *ProductRepo
}

func (service *ProductService) GetAllProducts() ([]Product, error) {
	return service.Repo.GetAllProducts()
}

func (service *ProductService) FilterProducts(name, productType string, minPrice, maxPrice, minRating float64) ([]Product, error) {
	filter := ProductFilter{
		Name:        name,
		ProductType: productType,
		MinPrice:    &minPrice,
		MaxPrice:    &maxPrice,
		MinRating:   &minRating,
	}

	return service.Repo.FilterProducts(filter)
}

func (service *ProductService) CreateProduct(product *Product) error {
	return service.Repo.CreateProduct(product)
}
