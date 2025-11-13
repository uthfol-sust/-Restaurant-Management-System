package models

type OrderDetail struct {
	OrderDetailID int64   `json:"order_detail_id"` // Primary Key
	OrderID       int64   `json:"order_id"`        // Foreign Key (references orders)
	ProductID     int   `json:"product_id"`      // Foreign Key (references products)
	Quantity      int     `json:"quantity"`        // Quantity of product
	Price         float64 `json:"price"`           // Price per unit
	Subtotal      float64 `json:"subtotal"`        // Quantity * Price
}
