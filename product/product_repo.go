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

func (repo *ProductRepo) FilterProducts (filter ProductFilter) ([]Product, error){
	var products []Product

	query := repo.DB


	if filter.Name != "" {
		query = query.Where("LOWER(product_type) LIKE LOWER(?)", "%"+filter.Name+"%")
	}
	if filter.MinPrice != nil {
		query = query.Where("price >= ?", *filter.MinPrice)
	}
	if filter.MaxPrice != nil {
		query = query.Where("price <= ?", *filter.MaxPrice)
	}
	if filter.ProductType != "" {
		query = query.Where("LOWER(product_type) = LOWER(?)", filter.ProductType)
	}
	if filter.MinRating != nil{
		query = query.Where("rating >= ?", filter.MinRating)
	}
	if err := query.Find(&products).Error; err != nil{
		return nil, err
	}

	return products, nil
}

func (repo *ProductRepo) CreateProduct(product *Product) error{
	return repo.DB.Create(product).Error
}