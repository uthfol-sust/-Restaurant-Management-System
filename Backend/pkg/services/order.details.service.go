package services

import (
	"errors"
	"restaurant-system/pkg/models"
	"restaurant-system/pkg/repositories"
)

type OrderDetailService interface {
	CreateOrderDetail(detail *models.OrderDetail) (*models.OrderDetail, error)
	GetByOrderID(orderID int64) ([]models.OrderDetail, error)
	DeleteOrderDetail(id int64) error
}

type orderDetailService struct {
	orderDetails repositories.OrderDetailsRepo
}

func NewOrderDetailService(service repositories.OrderDetailsRepo) OrderDetailService {
	return &orderDetailService{orderDetails: service}
}

func (s *orderDetailService) CreateOrderDetail(detail *models.OrderDetail) (*models.OrderDetail, error) {
	newDetails, err := s.orderDetails.CreateOrderDetail(detail)

	if err != nil {
		return nil, errors.New("service:failed to Create")
	}

	return newDetails, nil
}
func (s *orderDetailService) GetByOrderID(orderID int64) ([]models.OrderDetail, error) {
	details, err := s.orderDetails.GetByOrderID(orderID)
	if err != nil {
		return nil, errors.New("service:Not Found")
	}

	return details, nil
}
func (s *orderDetailService) DeleteOrderDetail(order_details_ID int64) error {

	items, err := s.orderDetails.GetByOrderDetailsID(order_details_ID)
	if err != nil {
		return err
	}
	if items == nil {
		return errors.New("service: order detail not found")
	}

	err = s.orderDetails.DeleteOrderDetail(order_details_ID)

	if err != nil {
		return  err
	}

	return nil
}
