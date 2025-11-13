package services

import (
	"errors"
	"restaurant-system/pkg/models"
	"restaurant-system/pkg/repositories"
)

type SupplierService interface {
	CreateSupplier(supplier *models.Supplier) (*models.Supplier, error)
	GetAllSupplier() ([]models.Supplier, error)
	GetBySupplierID(id int64) (*models.Supplier, error)
	UpdateSupplier(id int64, supplier *models.Supplier) (*models.Supplier, error)
	DeleteSupplier(id int64) error
}

type supplierService struct {
	supplierRepo repositories.Suppliers
}

func NewSupplierService(repo repositories.Suppliers) SupplierService {
	return &supplierService{supplierRepo: repo}
}

func (s *supplierService) CreateSupplier(supplier *models.Supplier) (*models.Supplier, error) {
	createdSupplier, err := s.supplierRepo.Create(supplier)
	if err != nil {
		return nil, err
	}
	return createdSupplier, nil
}

func (s *supplierService) GetAllSupplier() ([]models.Supplier, error) {
	suppliers, err := s.supplierRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return suppliers, nil
}

func (s *supplierService) GetBySupplierID(id int64) (*models.Supplier, error) {
	supplier, err := s.supplierRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return supplier, nil
}

func (s *supplierService) UpdateSupplier(id int64, supplier *models.Supplier) (*models.Supplier, error) {
	supplier.ID = int64(id)

	supp, err := s.supplierRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if supplier.Name != "" {
		supp.Name = supplier.Name
	}
	if supplier.Email != "" {
		supp.Address = supplier.Address
	}
	if supplier.Email != "" {
		supp.Email = supplier.Email
	}
	if supplier.ContactNo != "" {
		supp.ContactNo = supplier.ContactNo
	}

	updatedSupplier, err := s.supplierRepo.Update(supp)
	if err != nil {
		return nil, err
	}
	return updatedSupplier, nil
}

func (s *supplierService) DeleteSupplier(id int64) error {
	existing, err := s.supplierRepo.GetByID(id)
	if err != nil {
		return err
	}
	if existing == nil {
		return errors.New("service: cannot delete, supplier not found")
	}
	return s.supplierRepo.Delete(id)
}
