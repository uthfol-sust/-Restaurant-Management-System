package models

import "time"

type Payment struct {
	PaymentID     int64   `json:"payment_id"`     // Primary Key
	OrderID       int64   `json:"order_id"`       // Foreign Key (references orders)
	PaymentMethod string  `json:"payment_method"` // e.g., "Cash", "Card", "Online"
	PaymentStatus string  `json:"payment_status"` // e.g., "Paid", "Pending", "Failed"
	AmountPaid    float64 `json:"amount_paid"`    // Amount paid by customer
	PaymentDate   time.Time  `json:"payment_date"`   // Date of payment (YYYY-MM-DD)
}
