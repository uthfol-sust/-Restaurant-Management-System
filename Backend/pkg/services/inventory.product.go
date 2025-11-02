package services

import (
	"errors"
	"restaurant-system/pkg/models"
	"restaurant-system/pkg/repositories"
)

type InventoryProductService interface {
	CreateInventoryProduct(item *models.InventoryProduct) (*models.InventoryProduct, error)
	GetAllInventoryProduct() ([]models.InventoryProduct, error)
	GetInventoryProductByID(productID int) (*models.InventoryProduct, error)
	UpdateInventoryProduct(item *models.InventoryProduct) (*models.InventoryProduct, error)
	DeleteInventoryProduct(productID, inventoryID int) error
}

type inventoryProductService struct {
	inventory_product repositories.InventoryProductRepo
}

func NewInventoryProductService(inventoryProductRepo repositories.InventoryProductRepo) InventoryProductService {
	return &inventoryProductService{inventory_product: inventoryProductRepo}
}

func (ip *inventoryProductService) CreateInventoryProduct(item *models.InventoryProduct) (*models.InventoryProduct, error) {

	result, err := ip.inventory_product.CreateInventoryProduct(item)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (ip *inventoryProductService) GetAllInventoryProduct() ([]models.InventoryProduct, error) {
	items, err := ip.inventory_product.GetAllInventoryProduct()
	if err != nil {
		return nil, err
	}

	return *items, nil
}

func (ip *inventoryProductService) GetInventoryProductByID(productID int) (*models.InventoryProduct, error) {
	item, err := ip.inventory_product.GetInventoryProductByID(productID)
	if err != nil {
		return nil, err
	}

	if item == nil {
		return nil, errors.New("service: inventory product not found")
	}

	return item, nil
}


func (ip *inventoryProductService) UpdateInventoryProduct(item *models.InventoryProduct) (*models.InventoryProduct, error) {
	existing, err := ip.inventory_product.GetInventoryProductByID(item.ProductID)
	if err != nil {
		return nil, err
	}
	if existing == nil {
		return nil, errors.New("service: inventory product not found")
	}

	updatedItem, err := ip.inventory_product.UpdateInventoryProduct(item)
	if err != nil {
		return nil,  err
	}

	return updatedItem, nil
}



func (ip *inventoryProductService) DeleteInventoryProduct(productID, inventoryID int) error {
	existing, err := ip.inventory_product.GetInventoryProductByID(productID)
	if err != nil {
		return  err
	}
	if existing == nil {
		return errors.New("service: cannot delete, inventory product not found")
	}

	err = ip.inventory_product.DeleteInventoryProduct(productID, inventoryID)
	if err != nil {
		return  err
	}

	return nil
}
