package services

import (
	"errors"
	"restaurant-system/pkg/models"
	"restaurant-system/pkg/repositories"
	"time"
)

type PaymentService interface {
	CreatePayment(payment *models.Payment) (*models.Payment, error)
	GetAllPayments() ([]models.Payment, error)
	GetPaymentByID(id int64) (*models.Payment, error)
	UpdatePayment(payment *models.Payment) (*models.Payment, error)
	DeletePayment(id int64) error
}

type paymentService struct {
	paymentRepo repositories.PaymentsRepo
}

func NewPaymentService(payment repositories.PaymentsRepo) PaymentService {
	return &paymentService{paymentRepo: payment}
}

func (s *paymentService) CreatePayment(payment *models.Payment) (*models.Payment, error) {
	newPayment, err := s.paymentRepo.CreatePayment(payment)
	if err != nil {
		return nil, err
	}

	return newPayment, nil
}
func (s *paymentService) GetAllPayments() ([]models.Payment, error) {
	list, err := s.paymentRepo.GetAllPayments()
	if err != nil {
		return nil, err
	}

	return list, nil
}
func (s *paymentService) GetPaymentByID(id int64) (*models.Payment, error) {
	payment, err := s.paymentRepo.GetPaymentByID(id)
	if err != nil {
		return nil, err
	}

	return payment, nil
}
func (s *paymentService) UpdatePayment(payment *models.Payment) (*models.Payment, error) {
	newPayment, err := s.paymentRepo.GetPaymentByID(payment.PaymentID)
	if err != nil {
		return nil, errors.New("Service:Unvaild ID")
	}
    
	if payment.AmountPaid >=0 {
		newPayment.AmountPaid = payment.AmountPaid
	}
	if payment.PaymentStatus !=""{
		newPayment.PaymentStatus = payment.PaymentStatus
	}
	if payment.PaymentMethod !=""{
		newPayment.PaymentMethod = payment.PaymentMethod
	}
	if payment.PaymentDate.IsZero() {
		newPayment.PaymentDate = time.Now()
	}

	return s.paymentRepo.UpdatePayment(newPayment)
}
func (s *paymentService) DeletePayment(id int64) error {
	_, err := s.paymentRepo.GetPaymentByID(id)
	if err != nil {
		return errors.New("Service:Unvaild ID")
	}

	err = s.paymentRepo.DeletePayment(id)

	if err != nil {
		return errors.New("Service:Failed to Delete Payment")
	}

	return nil
}
