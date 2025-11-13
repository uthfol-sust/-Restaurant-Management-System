package controllers

import (
	"net/http"
	"restaurant-system/pkg/models"
	"restaurant-system/pkg/services"
	"restaurant-system/pkg/utils"
	"strconv"
)

type PaymentController interface {
	CreatePayment(w http.ResponseWriter, r *http.Request)
	GetAllPayments(w http.ResponseWriter, r *http.Request)
	GetPaymentByID(w http.ResponseWriter, r *http.Request)
	UpdatePayment(w http.ResponseWriter, r *http.Request)
	DeletePayment(w http.ResponseWriter, r *http.Request)
}

type paymentController struct {
	paymentService services.PaymentService
}

func NewPaymentController(s services.PaymentService) PaymentController {
	return &paymentController{paymentService: s}
}

func (c *paymentController) CreatePayment(w http.ResponseWriter, r *http.Request) {
	payment := &models.Payment{}
	utils.ParseBody(r, payment)

	createdPayment, err := c.paymentService.CreatePayment(payment)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	utils.HTTPResponse(w, 201, createdPayment)
}

func (c *paymentController) GetAllPayments(w http.ResponseWriter, r *http.Request) {
	payments, err := c.paymentService.GetAllPayments()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	utils.HTTPResponse(w, 200, payments)
}

func (c *paymentController) GetPaymentByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid payment ID", http.StatusBadRequest)
		return
	}

	payment, err := c.paymentService.GetPaymentByID(int64(id))
	if err != nil {
		http.Error(w, "Payment not found", 404)
		return
	}

	utils.HTTPResponse(w, 200, payment)
}

func (c *paymentController) UpdatePayment(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid payment ID", http.StatusBadRequest)
		return
	}

	payment := &models.Payment{}
	utils.ParseBody(r, payment)
	payment.PaymentID = int64(id)

	updatedPayment, err := c.paymentService.UpdatePayment(payment)
	if err != nil {
		http.Error(w, "Failed to update payment",500)
		return
	}

	utils.HTTPResponse(w, 200, updatedPayment)
}

func (c *paymentController) DeletePayment(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid payment ID", http.StatusBadRequest)
		return
	}

	err = c.paymentService.DeletePayment(int64(id))
	if err != nil {
		http.Error(w, "Failed to delete payment", 500)
		return
	}

	utils.HTTPResponse(w, http.StatusNoContent, "Payments Record Deleted")
}
