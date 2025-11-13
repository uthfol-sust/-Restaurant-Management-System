package repositories

import (
	"database/sql"
	"restaurant-system/pkg/models"
)

type PurchasesRepo interface {
	CreatePurchase(purchase *models.Purchase) (*models.Purchase, error)
	GetAllPurchases() ([]models.Purchase, error)
	GetPurchaseByID(id int64) (*models.Purchase, error)
	UpdatePurchase(purchase *models.Purchase) (*models.Purchase, error)
	DeletePurchase(id int64) error
}

type purchaseRepo struct {
	db *sql.DB
}

func NewPurchaseRepo(db *sql.DB) PurchasesRepo {
	return &purchaseRepo{db: db}
}

func (r *purchaseRepo) CreatePurchase(p *models.Purchase) (*models.Purchase, error) {
	query := `INSERT INTO purchases (purchase_id, supplier_id, inventory_id, quantity_purchased, purchase_date, total_cost)
			  VALUES (?, ?, ?, ?, ?, ?)`
	_, err := r.db.Exec(query,p.PurchaseID, p.SupplierID, p.InventoryID, p.QuantityPurchased, p.PurchaseDate, p.TotalCost)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (r *purchaseRepo) GetAllPurchases() ([]models.Purchase, error) {
	rows, err := r.db.Query(`SELECT * FROM purchases`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var purchases []models.Purchase
	for rows.Next() {
		p :=&models.Purchase{}
		if err := rows.Scan(&p.PurchaseID, &p.SupplierID, &p.InventoryID, &p.QuantityPurchased, &p.PurchaseDate, &p.TotalCost); err != nil {
			return nil, err
		}
		purchases = append(purchases, *p)
	}
	return purchases, nil
}

func (r *purchaseRepo) GetPurchaseByID(id int64) (*models.Purchase, error) {
	query := `SELECT * FROM purchases WHERE purchase_id = ?`
	row := r.db.QueryRow(query, id)
	p :=&models.Purchase{}
	err := row.Scan(&p.PurchaseID, &p.SupplierID, &p.InventoryID, &p.QuantityPurchased, &p.PurchaseDate, &p.TotalCost)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (r *purchaseRepo) UpdatePurchase(p *models.Purchase) (*models.Purchase, error) {
	query := `UPDATE purchases SET supplier_id=?, inventory_id=?, quantity_purchased=?, purchase_date=?, total_cost=? WHERE purchase_id=?`
	_, err := r.db.Exec(query, p.SupplierID, p.InventoryID, p.QuantityPurchased, p.PurchaseDate, p.TotalCost, p.PurchaseID)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (r *purchaseRepo) DeletePurchase(id int64) error {
	_, err := r.db.Exec(`DELETE FROM purchases WHERE purchase_id=?`, id)
	return err
}
