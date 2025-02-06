package product

import "gorm.io/gorm"

type ProductRepo struct {
	DB *gorm.DB
}

func (repo *ProductRepo) GetAllProducts() ([]Product, error){
	var products []Product

	if err := repo.DB.Find(&products).Error; err != nil{
		return nil, err
	}

	return products, nil
}

func (repo *ProductRepo) GetProductByName (name string) ([]Product, error){
	var products []Product
	if err := repo.DB.Where("name ILIKE ?", "%"+name+"%").Find(&products).Error; err != nil{
		return nil, err
	}

	return products, nil
}

func (repo *ProductRepo) CreateProduct(product *Product) error{
	return repo.DB.Create(product).Error
}