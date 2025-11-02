package models

type InventoryProduct struct {
	ProductID        int `json:"product_id"`        
	InventoryID      int `json:"id"`      
	QuantityRequired int `json:"quantity_required"`
}
