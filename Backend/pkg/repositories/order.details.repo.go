package repositories

import (
	"database/sql"
	"restaurant-system/pkg/models"
)

type OrderDetailsRepo interface {
	CreateOrderDetail(detail *models.OrderDetail) (*models.OrderDetail, error)
	GetByOrderID(orderID int64) ([]models.OrderDetail, error)
	GetByOrderDetailsID(orderDetailsID int64) (*models.OrderDetail, error)
	DeleteOrderDetail(id int64) error
}

type orderDetailRepo struct {
	db *sql.DB
}

func NewOrderDetailRepo(db *sql.DB) OrderDetailsRepo {
	return &orderDetailRepo{db: db}
}

func (r *orderDetailRepo) CreateOrderDetail(d *models.OrderDetail) (*models.OrderDetail, error) {
	query := `INSERT INTO order_details (order_detail_id, order_id, product_id, quantity, price, subtotal)
			  VALUES (?, ?, ?, ?, ?, ?)`
	_, err := r.db.Exec(query, d.OrderDetailID, d.OrderID, d.ProductID, d.Quantity, d.Price, d.Subtotal)
	if err != nil {
		return nil, err
	}
	return d, nil
}

func (r *orderDetailRepo) GetByOrderID(orderID int64) ([]models.OrderDetail, error) {
	rows, err := r.db.Query(`SELECT * FROM order_details WHERE order_id=?`, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var details []models.OrderDetail
	for rows.Next() {
		var d models.OrderDetail
		if err := rows.Scan(&d.OrderDetailID, &d.OrderID, &d.ProductID, &d.Quantity, &d.Price, &d.Subtotal); err != nil {
			return nil, err
		}
		details = append(details, d)
	}
	return details, nil
}

func (r *orderDetailRepo) GetByOrderDetailsID(orderDetailsID int64) (*models.OrderDetail, error) {
	var d models.OrderDetail
	err := r.db.QueryRow(`SELECT order_detail_id, order_id, product_id, quantity, price, subtotal FROM order_details WHERE order_detail_id=?`, orderDetailsID).
		Scan(&d.OrderDetailID, &d.OrderID, &d.ProductID, &d.Quantity, &d.Price, &d.Subtotal)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &d, nil
}

func (r *orderDetailRepo) DeleteOrderDetail(id int64) error {
	_, err := r.db.Exec(`DELETE FROM order_details WHERE order_detail_id=?`, id)
	return err
}
