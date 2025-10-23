package services

import (
	"errors"
	"restaurant-system/pkg/models"
	"restaurant-system/pkg/repositories"
)

type ProductService interface {
	CreateProduct(product *models.Product) (*models.Product, error)
	GetAllProduct() ([]models.Product, error)
	GetProductById(id int) (*models.Product, error)
	UpdateProduct(id int, product *models.Product) (*models.Product, error)
	DeleteProduct(id int) error
}

type productServices struct {
	productRepo repositories.ProductRepository
}


func NewProductService(productRepo repositories.ProductRepository) ProductService {
	return &productServices{productRepo: productRepo}
}

func (s *productServices) CreateProduct(product *models.Product) (*models.Product, error) {
	return s.productRepo.CreateProduct(product)
}

func (s *productServices) GetAllProduct() ([]models.Product, error) {
	return s.productRepo.GetAllProduct()
}

func (s *productServices) GetProductById(id int) (*models.Product, error) {
	return s.productRepo.GetProductById(id)
}

func (s *productServices) UpdateProduct(id int, product *models.Product) (*models.Product, error) {
	curProduct, err := s.productRepo.GetProductById(id)
	if err != nil {
		return nil, errors.New("product not found to updated")
	}

	if product.ProductName != "" {
		curProduct.ProductName = product.ProductName
	}
	if product.Category != "" {
		curProduct.Category = product.Category
	}
	if product.Price > 0.0 {
		curProduct.Price = product.Price
	}
	if product.AvailabilityStatus!=""{
		curProduct.AvailabilityStatus=product.AvailabilityStatus
	}

	return s.productRepo.UpdateProduct(id, curProduct)
}

func (s *productServices) DeleteProduct(id int) error {
	return s.productRepo.DeleteProduct(id)
}
