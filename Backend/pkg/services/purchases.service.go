package services

import (
	"restaurant-system/pkg/models"
	"restaurant-system/pkg/repositories"
)

type PurchaseService interface {
	CreatePurchase(purchase *models.Purchase) (*models.Purchase, error)
	GetAllPurchases() ([]models.Purchase, error)
	GetPurchaseByID(id int64) (*models.Purchase, error)
	UpdatePurchase(purchase *models.Purchase, id int64) (*models.Purchase, error)
	DeletePurchase(id int64) error
}

type purchaseService struct {
	purchaseRepo repositories.PurchasesRepo
}

func NewPurchaseService(Repo repositories.PurchasesRepo) PurchaseService {
	return &purchaseService{purchaseRepo: Repo}
}

func (s *purchaseService) CreatePurchase(purchase *models.Purchase) (*models.Purchase, error) {
	newPurchase, err := s.purchaseRepo.CreatePurchase(purchase)
	if err != nil {
		return nil, err
	}

	return newPurchase, nil
}

func (s *purchaseService) GetAllPurchases() ([]models.Purchase, error) {
	allPurchases, err := s.purchaseRepo.GetAllPurchases()
	if err != nil {
		return nil, err
	}

	return allPurchases, nil
}
func (s *purchaseService) GetPurchaseByID(id int64) (*models.Purchase, error) {
	purchase, err := s.purchaseRepo.GetPurchaseByID(id)
	if err != nil {
		return nil, err
	}

	return purchase, nil
}
func (s *purchaseService) UpdatePurchase(purchase *models.Purchase, id int64) (*models.Purchase, error) {
	newPurchase, err := s.purchaseRepo.GetPurchaseByID(id)
	if err != nil {
		return nil, err
	}

	if purchase.QuantityPurchased >= 0 {
		newPurchase.QuantityPurchased = purchase.QuantityPurchased
	}
	if purchase.TotalCost >= 0 {
		newPurchase.TotalCost = purchase.TotalCost
	}
	if purchase.PurchaseDate != "" {
		newPurchase.PurchaseDate = purchase.PurchaseDate
	}

	result, err := s.purchaseRepo.UpdatePurchase(newPurchase)
	if err != nil {
		return nil, err
	}

	return result, nil
}
func (s *purchaseService) DeletePurchase(id int64) error {
	_, err := s.purchaseRepo.GetPurchaseByID(id)
	if err != nil {
		return err
	}

	err = s.purchaseRepo.DeletePurchase(id)
	if err != nil {
		return err
	}

	return nil
}
