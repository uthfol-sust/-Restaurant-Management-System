package controllers

import (
	"net/http"
	"restaurant-system/pkg/models"
	"restaurant-system/pkg/services"
	"restaurant-system/pkg/utils"
	"strconv"
)

type PurchaseController interface {
	CreatePurchase(w http.ResponseWriter, r *http.Request)
	GetAllPurchases(w http.ResponseWriter, r *http.Request)
	GetPurchaseByID(w http.ResponseWriter, r *http.Request)
	UpdatePurchase(w http.ResponseWriter, r *http.Request)
	DeletePurchase(w http.ResponseWriter, r *http.Request)
}

type purchaseController struct {
	purchaseService services.PurchaseService
}

func NewPurchaseController(s services.PurchaseService) PurchaseController {
	return &purchaseController{purchaseService: s}
}

func (c *purchaseController) CreatePurchase(w http.ResponseWriter, r *http.Request) {
	payment := &models.Purchase{}
	utils.ParseBody(r, payment)

	respose, err := c.purchaseService.CreatePurchase(payment)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	utils.HTTPResponse(w, 201, respose)
}

func (c *purchaseController) GetAllPurchases(w http.ResponseWriter, r *http.Request) {
	paymentList, err := c.purchaseService.GetAllPurchases()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	utils.HTTPResponse(w, 200, paymentList)
}

func (c *purchaseController) GetPurchaseByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invaild ID", http.StatusBadRequest)
		return
	}

	payment, err := c.purchaseService.GetPurchaseByID(int64(id))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	utils.HTTPResponse(w, 200, payment)
}

func (c *purchaseController) UpdatePurchase(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.PathValue("id"))
	payment := &models.Purchase{}
	utils.ParseBody(r, payment)

	updated, err := c.purchaseService.UpdatePurchase(payment, int64(id))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	utils.HTTPResponse(w, 200, updated)
}

func (c *purchaseController) DeletePurchase(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invaild ID", http.StatusBadRequest)
		return
	}
	err = c.purchaseService.DeletePurchase(int64(id))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	utils.HTTPResponse(w, 200, "Purchases Deleted")
}
