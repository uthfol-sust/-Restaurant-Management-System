package repositories

import (
	"database/sql"
	"errors"
	"restaurant-system/pkg/models"
)

type InventoryRepository interface {
	CreateInventory(item *models.Inventory) (*models.Inventory, error)
	GetAllInventory() (*[]models.Inventory, error)
	GetInventoryByID(id int64) (*models.Inventory, error)
	UpdateInventory(item *models.Inventory) (*models.Inventory, error)
	DeleteInventory(id int64) error
}

type inventoryRepository struct {
	db *sql.DB
}

func NewInventoryRepository(db *sql.DB) InventoryRepository {
	return &inventoryRepository{db: db}
}

func (it *inventoryRepository) CreateInventory(item *models.Inventory) (*models.Inventory, error) {
	query := `INSERT INTO inventory (id, name, stock, unit, level) 
	          VALUES (?, ?, ?, ?, ?)`

	result, err := it.db.Exec(query,item.InventoryID, item.ItemName, item.QuantityInStock, item.Unit, item.ReorderLevel)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	item.InventoryID = id
	return item, nil
}

func (it *inventoryRepository) GetAllInventory() (*[]models.Inventory, error) {
	query := `SELECT id, name, stock, unit, level, last_updated FROM inventory`

	rows, err := it.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var inventories []models.Inventory

	for rows.Next() {
		inv := &models.Inventory{}
		err := rows.Scan(&inv.InventoryID, &inv.ItemName, &inv.QuantityInStock, &inv.Unit, &inv.ReorderLevel, &inv.LastUpdated)
		if err != nil {
			return nil, err
		}
		inventories = append(inventories, *inv)
	}

	return &inventories, nil
}

func (it *inventoryRepository) GetInventoryByID(id int64) (*models.Inventory, error) {
	query := `SELECT id, name, stock, unit, level, last_updated 
	          FROM inventory WHERE id = ?`

	inv := &models.Inventory{}
	err := it.db.QueryRow(query, id).Scan(
		&inv.InventoryID, &inv.ItemName, &inv.QuantityInStock, &inv.Unit, &inv.ReorderLevel, &inv.LastUpdated,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("inventory item not found")
		}
		return nil, err
	}

	return inv, nil
}

func (it *inventoryRepository) UpdateInventory(item *models.Inventory) (*models.Inventory, error) {
	query := `UPDATE inventory 
	          SET name = ?, stock = ?, unit = ?, level = ?, last_updated = ? 
	          WHERE id = ?`

	_, err := it.db.Exec(query, item.ItemName, item.QuantityInStock, item.Unit, item.ReorderLevel, item.LastUpdated, item.InventoryID)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (it *inventoryRepository) DeleteInventory(id int64) error {
	query := `DELETE FROM inventory WHERE id = ?`

	result, err := it.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no inventory record found to delete")
	}

	return nil
}
