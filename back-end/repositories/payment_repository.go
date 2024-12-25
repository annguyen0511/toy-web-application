package repositories

import (
	"backend/models"
	"database/sql"
	"time"
)

type PaymentRepository struct {
	DB *sql.DB
}

func NewPaymentRepository(db *sql.DB) *PaymentRepository {
	return &PaymentRepository{DB: db}
}

// Thêm thanh toán
func (r *PaymentRepository) AddPayment(payment models.Payment) error {
	_, err := r.DB.Exec("INSERT INTO Payments (OrderID, Amount, PaymentDate, Status) VALUES (?, ?, ?, ?)",
		payment.OrderID, payment.Amount, payment.PaymentDate, payment.Status)
	return err
}

// Sửa thanh toán
func (r *PaymentRepository) UpdatePayment(payment models.Payment) error {
	_, err := r.DB.Exec("UPDATE Payments SET Amount = ?, PaymentDate = ?, Status = ? WHERE PaymentID = ?",
		payment.Amount, payment.PaymentDate, payment.Status, payment.PaymentID)
	return err
}

// Xóa thanh toán
func (r *PaymentRepository) DeletePayment(paymentID int) error {
	_, err := r.DB.Exec("DELETE FROM Payments WHERE PaymentID = ?", paymentID)
	return err
}

// Lấy tất cả thanh toán
func (r *PaymentRepository) GetAllPayments() ([]models.Payment, error) {
	rows, err := r.DB.Query("SELECT PaymentID, OrderID, Amount, PaymentDate, Status FROM Payments")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var payments []models.Payment
	for rows.Next() {
		var payment models.Payment
		var paymentDate string // Temporary variable to hold the PaymentDate value
		if err := rows.Scan(&payment.PaymentID, &payment.OrderID, &payment.Amount, &paymentDate, &payment.Status); err != nil {
			return nil, err
		}
		// Convert paymentDate string to time.Time
		if parsedTime, err := time.Parse("2006-01-02 15:04:05", paymentDate); err == nil {
			payment.PaymentDate = parsedTime
		} else {
			println("Error parsing PaymentDate:", err.Error()) // Print error
			return nil, err
		}
		payments = append(payments, payment)
	}
	return payments, nil
}

// Lấy thanh toán theo ID
func (r *PaymentRepository) GetPaymentByID(paymentID int) (models.Payment, error) {
	var payment models.Payment
	err := r.DB.QueryRow("SELECT PaymentID, OrderID, Amount, PaymentDate, Status FROM Payments WHERE PaymentID = ?", paymentID).Scan(&payment.PaymentID, &payment.OrderID, &payment.Amount, &payment.PaymentDate, &payment.Status)
	return payment, err
}
