package services

import (
	"errors"
	"restaurant-system/pkg/models"
	"restaurant-system/pkg/repositories"
)

type OrderService interface {
	CreateOrder(order *models.Order) (*models.Order, error)
	GetAllOrders() ([]models.Order, error)
	GetOrderByID(id int64) (*models.Order, error)
	UpdateOrder(order *models.Order) (*models.Order, error)
	DeleteOrder(id int64) error
}

type orderService struct {
	orderRepo repositories.OrdersRepo
}

func NewOrderService(Repo repositories.OrdersRepo) OrderService {
	return &orderService{orderRepo: Repo}
}

func (s *orderService) CreateOrder(newOrder *models.Order) (*models.Order, error) {
	ordered, err := s.orderRepo.CreateOrder(newOrder)

	if err != nil {
		return nil, err
	}

	return ordered, nil
}
func (s *orderService) GetAllOrders() ([]models.Order, error) {
	list, err := s.orderRepo.GetAllOrders()

	if err != nil {
		return nil, errors.New("service:failed to load orders list")
	}

	return list, nil
}
func (s *orderService) GetOrderByID(id int64) (*models.Order, error) {
	order, err := s.orderRepo.GetOrderByID(id)
	if err != nil {
		return nil, errors.New("service:Order not found")
	}

	return order, nil
}
func (s *orderService) UpdateOrder(order *models.Order) (*models.Order, error) {
	updateOrder, err := s.orderRepo.GetOrderByID(order.OrderID)
	if err != nil {
		return nil, err
	}

	if order.Status != "" {
		updateOrder.Status = order.Status
	}
	if !order.OrderTime.IsZero() {
		updateOrder.OrderTime = order.OrderTime
	}
	if order.TotalAmount >= 0 {
		updateOrder.TotalAmount = order.TotalAmount
	}

	_, err = s.orderRepo.UpdateOrder(updateOrder)
	if err != nil {
		return nil, err
	}

	return updateOrder, nil
}
func (s *orderService) DeleteOrder(id int64) error {
	_, err := s.orderRepo.GetOrderByID(id)
	if err != nil {
		return err
	}

	err = s.orderRepo.DeleteOrder(id)
	if err != nil {
		return err
	}

	return nil
}
