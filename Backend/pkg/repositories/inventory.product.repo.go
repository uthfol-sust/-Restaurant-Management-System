package repositories

import (
	"database/sql"
	"fmt"
	"restaurant-system/pkg/models"
)

type InventoryProductRepo interface {
	CreateInventoryProduct(item *models.InventoryProduct) (*models.InventoryProduct, error)
	GetAllInventoryProduct() (*[]models.InventoryProduct, error)
	GetInventoryProductByID(productID int) (*models.InventoryProduct, error)
	UpdateInventoryProduct(item *models.InventoryProduct) (*models.InventoryProduct, error)
	DeleteInventoryProduct(productID, inventoryID int) error
}

type inventoryProductRepo struct {
	db *sql.DB
}

func NewInventoryProductRepo(db *sql.DB) InventoryProductRepo {
	return &inventoryProductRepo{db: db}
}


func (r *inventoryProductRepo) CreateInventoryProduct(item *models.InventoryProduct) (*models.InventoryProduct, error) {
	query := `INSERT INTO inventory_product (product_id, id, quantity_required)
			  VALUES (?, ?, ?)`

	_, err := r.db.Exec(query, item.ProductID, item.InventoryID, item.QuantityRequired)
	if err != nil {
		return nil, fmt.Errorf("failed to insert inventory product: %v", err)
	}
	return item, nil
}


func (r *inventoryProductRepo) GetAllInventoryProduct() (*[]models.InventoryProduct, error) {
	query := `SELECT product_id, id, quantity_required FROM inventory_product`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch inventory products: %v", err)
	}
	defer rows.Close()

	var items []models.InventoryProduct
	for rows.Next() {
		var item models.InventoryProduct
		if err := rows.Scan(&item.ProductID, &item.InventoryID, &item.QuantityRequired); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return &items, nil
}


func (r *inventoryProductRepo) GetInventoryProductByID(productID int) (*models.InventoryProduct, error) {
	query := `SELECT product_id, id, quantity_required
			  FROM inventory_product
			  WHERE product_id = ?`

	var item models.InventoryProduct
	err := r.db.QueryRow(query, productID).Scan(
		&item.ProductID,
		&item.InventoryID,
		&item.QuantityRequired,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get inventory product: %v", err)
	}
	return &item, nil
}


func (r *inventoryProductRepo) UpdateInventoryProduct(item *models.InventoryProduct) (*models.InventoryProduct, error) {
	query := `UPDATE inventory_product
			  SET quantity_required = ?
			  WHERE product_id = ? AND id = ?`

	_, err := r.db.Exec(query, item.QuantityRequired, item.ProductID, item.InventoryID)
	if err != nil {
		return nil, fmt.Errorf("failed to update inventory product: %v", err)
	}
	return item, nil
}


func (r *inventoryProductRepo) DeleteInventoryProduct(productID, inventoryID int) error {
	query := `DELETE FROM inventory_product WHERE product_id = ? AND id = ?`
	_, err := r.db.Exec(query, productID, inventoryID)
	if err != nil {
		return fmt.Errorf("failed to delete inventory product: %v", err)
	}
	return nil
}
