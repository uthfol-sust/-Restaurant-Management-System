package models

type Purchase struct {
	PurchaseID        int64   `json:"purchase_id"`         // Primary Key
	SupplierID        int64   `json:"supplier_id"`         // Foreign Key (references suppliers)
	InventoryID       int64   `json:"inventory_id"`        // Foreign Key (references inventory)
	QuantityPurchased  float32     `json:"quantity_purchased"`  // Quantity purchased
	PurchaseDate      string  `json:"purchase_date"`       // Date of purchase (YYYY-MM-DD)
	TotalCost         float64 `json:"total_cost"`          // Total purchase cost
}
