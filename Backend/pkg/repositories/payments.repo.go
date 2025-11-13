package repositories

import (
	"database/sql"
	"restaurant-system/pkg/models"
	"time"
)

type PaymentsRepo interface {
	CreatePayment(payment *models.Payment) (*models.Payment, error)
	GetAllPayments() ([]models.Payment, error)
	GetPaymentByID(id int64) (*models.Payment, error)
	UpdatePayment(payment *models.Payment) (*models.Payment, error)
	DeletePayment(id int64) error
}

type paymentRepo struct {
	db *sql.DB
}

func NewPaymentRepo(db *sql.DB) PaymentsRepo {
	return &paymentRepo{db: db}
}

func (r *paymentRepo) CreatePayment(p *models.Payment) (*models.Payment, error) {
	query := `INSERT INTO payments (payment_id, order_id, payment_method, payment_status, amount_paid, payment_date)
			  VALUES (?, ?, ?, ?, ?, ?)`

	formatTime := p.PaymentDate.Format("2006-01-02 15:04:05")

	_, err := r.db.Exec(query, p.PaymentID, p.OrderID, p.PaymentMethod, p.PaymentStatus, p.AmountPaid, formatTime)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (r *paymentRepo) GetAllPayments() ([]models.Payment, error) {
	rows, err := r.db.Query(`SELECT payment_id, order_id, payment_method, payment_status, amount_paid, payment_date FROM payments`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var payments []models.Payment
	for rows.Next() {
		var formatTime string
		var p models.Payment
		if err := rows.Scan(&p.PaymentID, &p.OrderID, &p.PaymentMethod, &p.PaymentStatus, &p.AmountPaid, &formatTime); err != nil {
			return nil, err
		}

		p.PaymentDate, _ = time.Parse("2006-01-02 15:04:05", formatTime)

		payments = append(payments, p)
	}
	return payments, nil
}

func (r *paymentRepo) GetPaymentByID(id int64) (*models.Payment, error) {
	row := r.db.QueryRow(`SELECT * FROM payments WHERE payment_id=?`, id)
	var p models.Payment
	var formatTime string
	err := row.Scan(&p.PaymentID, &p.OrderID, &p.PaymentMethod, &p.PaymentStatus, &p.AmountPaid, &formatTime)
	if err != nil {
		return nil, err
	}

	p.PaymentDate, err = time.Parse("2006-01-02 15:04:05", formatTime)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *paymentRepo) UpdatePayment(p *models.Payment) (*models.Payment, error) {
	query := `UPDATE payments SET order_id=?, payment_method=?, payment_status=?, amount_paid=?, payment_date=? WHERE payment_id=?`

	formatTime := p.PaymentDate.Format("2006-01-02 15:04:05")

	_, err := r.db.Exec(query, p.OrderID, p.PaymentMethod, p.PaymentStatus, p.AmountPaid, formatTime, p.PaymentID)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (r *paymentRepo) DeletePayment(id int64) error {
	_, err := r.db.Exec(`DELETE FROM payments WHERE payment_id=?`, id)
	return err
}
