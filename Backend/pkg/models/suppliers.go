package models


type Supplier struct {
    ID         int64  `json:"supplier_id"` 
    Name       string `json:"name"`       
    ContactNo  string `json:"contact_no"` 
    Email      string `json:"email"`      
    Address    string `json:"address"`
}
