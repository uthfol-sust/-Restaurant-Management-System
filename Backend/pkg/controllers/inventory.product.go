package controllers

import (
	"net/http"
	"strconv"

	"restaurant-system/pkg/models"
	"restaurant-system/pkg/services"
	"restaurant-system/pkg/utils"
)

type InventoryProductController interface {
	CreateInventoryProduct(w http.ResponseWriter, r *http.Request)
	GetAllInventoryProduct(w http.ResponseWriter, r *http.Request)
	GetInventoryProductByID(w http.ResponseWriter, r *http.Request)
	UpdateInventoryProduct(w http.ResponseWriter, r *http.Request)
	DeleteInventoryProduct(w http.ResponseWriter, r *http.Request)
}

type inventoryProductController struct {
	service services.InventoryProductService
}

func NewInventoryProductController(service services.InventoryProductService) InventoryProductController {
	return &inventoryProductController{service: service}
}

func (c *inventoryProductController) CreateInventoryProduct(w http.ResponseWriter, r *http.Request) {
	item := &models.InventoryProduct{}
	utils.ParseBody(r, item)

	products, err := c.service.CreateInventoryProduct(item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.HTTPResponse(w, 201, products)
}

func (c *inventoryProductController) GetAllInventoryProduct(w http.ResponseWriter, r *http.Request) {
	items, err := c.service.GetAllInventoryProduct()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.HTTPResponse(w, 200, items)
}

func (c *inventoryProductController) GetInventoryProductByID(w http.ResponseWriter, r *http.Request) {
	productID, err1 := strconv.Atoi(r.PathValue("id"))

	if err1 != nil {
		http.Error(w, "Invalid ID parameters", http.StatusBadRequest)
		return
	}

	item, err := c.service.GetInventoryProductByID(productID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	utils.HTTPResponse(w, 200, item)
}

func (c *inventoryProductController) UpdateInventoryProduct(w http.ResponseWriter, r *http.Request) {
	item := &models.InventoryProduct{}
	utils.ParseBody(r, item)

	updatedItem, err := c.service.UpdateInventoryProduct(item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.HTTPResponse(w, 200, updatedItem)
}

func (c *inventoryProductController) DeleteInventoryProduct(w http.ResponseWriter, r *http.Request) {
	productID, err1 := strconv.Atoi(r.PathValue("product_id"))
	inventoryID, err2 := strconv.Atoi(r.PathValue("inventory_id"))

	if err1 != nil || err2 != nil {
		http.Error(w, "Invalid ID parameters", http.StatusBadRequest)
		return
	}

	if err := c.service.DeleteInventoryProduct(productID, inventoryID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.HTTPResponse(w, 200, "Item deleted")
}
