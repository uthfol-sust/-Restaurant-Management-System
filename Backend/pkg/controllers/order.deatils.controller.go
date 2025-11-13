package controllers

import (
	"net/http"
	"restaurant-system/pkg/models"
	"restaurant-system/pkg/services"
	"restaurant-system/pkg/utils"
	"strconv"
)

type OrderDetailController interface {
	CreateOrderDetail(w http.ResponseWriter, r *http.Request)
	GetByOrderID(w http.ResponseWriter, r *http.Request)
	DeleteOrderDetail(w http.ResponseWriter, r *http.Request)
}

type orderDetailsController struct {
	orderDetailService services.OrderDetailService
}

func NewOrderDetailController(s services.OrderDetailService) OrderDetailController {
	return &orderDetailsController{orderDetailService: s}
}

func (c *orderDetailsController) CreateOrderDetail(w http.ResponseWriter, r *http.Request) {
	orderDetail := &models.OrderDetail{}
	utils.ParseBody(r, orderDetail)

	createdDetail, err := c.orderDetailService.CreateOrderDetail(orderDetail)
	if err != nil {
		http.Error(w, "Failed to create order detail", 500)
		return
	}

	utils.HTTPResponse(w, 201, createdDetail)
}

func (c *orderDetailsController) GetByOrderID(w http.ResponseWriter, r *http.Request) {
	orderID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	details, err := c.orderDetailService.GetByOrderID(int64(orderID))
	if err != nil {
		http.Error(w, err.Error(),500)
		return
	}

	utils.HTTPResponse(w, 200, details)
}

func (c *orderDetailsController) DeleteOrderDetail(w http.ResponseWriter, r *http.Request) {
	order_details_ID, err := strconv.Atoi(r.PathValue("order_detail_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.orderDetailService.DeleteOrderDetail(int64(order_details_ID))
	if err != nil {
		http.Error(w, err.Error(),500)
		return
	}

	utils.HTTPResponse(w, 200, "Item is Deleted")
}
