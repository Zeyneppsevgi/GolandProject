package service

import (
	"product-app/domain"
	"product-app/persistence"
	"product-app/service/model"
)

type IProductService interface {
	Add(productCreate model.ProductCreate)
	DeleteById(productId int64) error
	GetById(productId int64) (domain.Product, error)
	UpdatePrice(productId int64, newPrice float32) error
	GetAllProducts() []domain.Product
	GetAllProductsByStore(storeName string) []domain.Product
}
type ProductService struct {
	productRepository persistence.IProductRepository
}

func NewProductService(productRepository persistence.IProductRepository) IProductService {
	return &ProductRepository{
		dbPool: dbPool,
	}
}

func (productService *ProductService) Add(productCreate model.ProductCreate) error {

}

func (productService *ProductService) DeleteById(productId int64) error {

}
func (productService *ProductService) GetById(productId int64) (domain.Product, error) {

}

func (productService *ProductService) UpdatePrice(productId int64, newPrice float32) error {

}

func (productService *ProductService) GetAllProducts() []domain.Product {

}

func (productService *ProductService) GetAllProductsByStore(storeName string) []domain.Product {
	
}
