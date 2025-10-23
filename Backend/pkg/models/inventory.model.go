package models

type Inventory struct {
    InventoryID     int64     `json:"id"`
    ItemName        string  `json:"name"`
    QuantityInStock float64 `json:"stock"`
    Unit            string  `json:"unit"`
    ReorderLevel    float64 `json:"level"`
    LastUpdated     string  `json:"last_updated"`
}
