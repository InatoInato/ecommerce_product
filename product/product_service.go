package product

type ProductService struct {
	Repo *ProductRepo
}

func (service *ProductService) GetAllProducts() ([]Product, error) {
	return service.Repo.GetAllProducts()
}

func (service *ProductService) SearchProductByName(name string) ([]Product, error) {
	return service.Repo.GetProductByName(name)
}


func (service *ProductService) CreateProduct(product *Product) error{
	return service.Repo.CreateProduct(product)
}

