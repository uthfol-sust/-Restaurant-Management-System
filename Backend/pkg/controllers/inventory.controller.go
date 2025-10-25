package controllers

import (
	"net/http"
	"restaurant-system/pkg/models"
	"restaurant-system/pkg/services"
	"restaurant-system/pkg/utils"
	"strconv"
)

type InventoryController interface {
	CreateInventory(w http.ResponseWriter, r *http.Request)
	GetAllInventory(w http.ResponseWriter, r *http.Request)
	GetInventoryByID(w http.ResponseWriter, r *http.Request)
	UpdateInventory(w http.ResponseWriter, r *http.Request)
	DeleteInventory(w http.ResponseWriter, r *http.Request)
}

type inventoryController struct {
	inventoryServ services.InventoryService
}

func NewInventoryController(services services.InventoryService) InventoryController {
	return &inventoryController{inventoryServ: services}
}

func (c *inventoryController) CreateInventory(w http.ResponseWriter, r *http.Request) {
	item := &models.Inventory{}
	utils.ParseBody(r, item)

	itemCreated, err := c.inventoryServ.CreateInventory(item)
	if err != nil {
		http.Error(w, "Failed to create inventory item", http.StatusInternalServerError)
		return
	}

	utils.HTTPResponse(w, 201, itemCreated)
}
func (c *inventoryController) GetAllInventory(w http.ResponseWriter, r *http.Request) {
	items, err := c.inventoryServ.GetAllInventory()
	if err != nil {
		http.Error(w, "Failed to fetch inventory items", http.StatusInternalServerError)
		return
	}

	utils.HTTPResponse(w, http.StatusOK, items)
}
func (c *inventoryController) GetInventoryByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid inventory ID", http.StatusBadRequest)
		return
	}

	item, err := c.inventoryServ.GetInventoryByID(int64(id))
	if err != nil {
		http.Error(w, "Inventory item not found", http.StatusNotFound)
		return
	}

	utils.HTTPResponse(w, http.StatusOK, item)
}
func (c *inventoryController) UpdateInventory(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid inventory ID", http.StatusBadRequest)
		return
	}

	item := &models.Inventory{}
	utils.ParseBody(r, item)

	updatedItem, err := c.inventoryServ.UpdateInventory(int64(id), item)
	if err != nil {
		http.Error(w, "Failed to update inventory item", http.StatusInternalServerError)
		return
	}

	utils.HTTPResponse(w, http.StatusOK, updatedItem)
}
func (c *inventoryController) DeleteInventory(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid inventory ID", http.StatusBadRequest)
		return
	}

	err = c.inventoryServ.DeleteInventory(int64(id))
	if err != nil {
		http.Error(w, "Inventory item not found or delete failed", http.StatusInternalServerError)
		return
	}

	utils.HTTPResponse(w, http.StatusOK, map[string]string{"message": "Deleted successfully"})

}
