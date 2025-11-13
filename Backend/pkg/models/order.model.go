package models

import "time"

type Order struct {
	OrderID     int64   `json:"order_id"`     // Primary Key
	WaiterID    int     `json:"waiter_id"`    // Foreign Key (references users)
	CustomerID  int     `json:"customer_id"`  // Foreign Key (references users)
	OrderTime   time.Time  `json:"order_time"`   // Timestamp (e.g., "2025-11-07T15:04:05Z")
	Status      string  `json:"status"`       // e.g., "Pending", "Completed", "Cancelled"
	TotalAmount float64 `json:"total_amount"` // Total order amount
}
