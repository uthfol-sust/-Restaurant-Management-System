package repositories

import (
	"database/sql"
	"restaurant-system/pkg/models"
)

type Suppliers interface {
	Create(supplier *models.Supplier) (*models.Supplier, error)
	GetAll() ([]models.Supplier, error)
	GetByID(id int64) (*models.Supplier, error)
	Update(supplier *models.Supplier) (*models.Supplier, error)
	Delete(id int64) error
}

type suppliers struct {
	db *sql.DB
}

func NewSupplierRepository(db *sql.DB) Suppliers {
	return &suppliers{db: db}
}

func (r *suppliers) Create(supplier *models.Supplier) (*models.Supplier, error) {
	query := `INSERT INTO suppliers (supplier_id, name, contact_no, email, address) VALUES (?, ?, ?, ?, ?)`
	result, err := r.db.Exec(query, supplier.ID, supplier.Name, supplier.ContactNo, supplier.Email, supplier.Address)
	if err != nil {
		return nil, err
	}
	supplier.ID , _ = result.LastInsertId()
	return supplier, nil
}

func (r *suppliers) GetAll() ([]models.Supplier, error) {
	query := `SELECT supplier_id, name, contact_no, email, address FROM suppliers`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var suppliers []models.Supplier
	for rows.Next() {
		 s := &models.Supplier{}
		if err := rows.Scan(&s.ID, &s.Name, &s.ContactNo, &s.Email, &s.Address); err != nil {
			return nil, err
		}
		suppliers = append(suppliers, *s)
	}
	return suppliers, nil
}

func (r *suppliers) GetByID(id int64) (*models.Supplier, error) {
	query := `SELECT supplier_id, name, contact_no, email, address FROM suppliers WHERE supplier_id = ?`
	row := r.db.QueryRow(query, id)

	s :=&models.Supplier{}
	if err := row.Scan(&s.ID, &s.Name, &s.ContactNo, &s.Email, &s.Address); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil 
		}
		return nil, err
	}
	return s, nil
}

func (r *suppliers) Update(supplier *models.Supplier) (*models.Supplier, error) {
	query := `UPDATE suppliers SET name = ?, contact_no = ?, email = ?, address = ? WHERE supplier_id = ?`
	_, err := r.db.Exec(query, supplier.Name, supplier.ContactNo, supplier.Email, supplier.Address, supplier.ID)
	if err != nil {
		return nil, err
	}
	return supplier, nil
}

func (r *suppliers) Delete(id int64) error {
	query := `DELETE FROM suppliers WHERE supplier_id = ?`
	_, err := r.db.Exec(query, id)
	return err
}
