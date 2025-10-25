package services

import (
	"errors"
	"restaurant-system/pkg/models"
	"restaurant-system/pkg/repositories"
)

type InventoryService interface {
	CreateInventory(item *models.Inventory) (*models.Inventory, error)
	GetAllInventory() (*[]models.Inventory, error)
	GetInventoryByID(id int64) (*models.Inventory, error)
	UpdateInventory(id int64,item *models.Inventory) (*models.Inventory, error)
	DeleteInventory(id int64) error
}

type inventoryService struct{
	inventoryRepo repositories.InventoryRepository
}


func NewInventoryService(repo repositories.InventoryRepository) InventoryService{
	return &inventoryService{inventoryRepo: repo}
}

func (s inventoryService) CreateInventory(item *models.Inventory) (*models.Inventory, error){
    return s.inventoryRepo.CreateInventory(item)
}
func (s inventoryService) GetAllInventory() (*[]models.Inventory, error){
  return s.inventoryRepo.GetAllInventory()
}
func (s inventoryService) GetInventoryByID(id int64) (*models.Inventory, error){
  return  s.inventoryRepo.GetInventoryByID(id)
}
func (s inventoryService) UpdateInventory(id int64, item *models.Inventory) (*models.Inventory, error){
  exitInventory , err:= s.inventoryRepo.GetInventoryByID(id)
  if err !=nil{
	return nil, errors.New("this inventory is not found to update")
  }

  if item.ItemName!=""{
	exitInventory.ItemName = item.ItemName
  }
  if item.Unit !=""{
	exitInventory.Unit=item.Unit
  }
  if item.ReorderLevel > 0{
	exitInventory.ReorderLevel=item.ReorderLevel
  }
  if item.QuantityInStock>0{
	exitInventory.QuantityInStock=item.QuantityInStock
  }
  if item.LastUpdated !=""{
	exitInventory.LastUpdated=item.LastUpdated
  }

  return s.inventoryRepo.UpdateInventory(exitInventory)
}
func (s inventoryService) DeleteInventory(id int64) error{
   _, err:= s.inventoryRepo.GetInventoryByID(id)
  if err !=nil{
	return errors.New("this inventory is not found to delete")
  }

  return s.inventoryRepo.DeleteInventory(id)
}
