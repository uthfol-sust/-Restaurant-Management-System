package models 

type Product struct {
    ProductID          int     `json:"product_id"`          
    ProductName        string  `json:"product_name"`
    Category           string  `json:"category"`
    Price              float64 `json:"price"`
    AvailabilityStatus string  `json:"availability_status"`
}