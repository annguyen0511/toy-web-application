package services

import (
	"backend/models"
	"backend/repositories"
)

type PaymentService struct {
	PaymentRepo *repositories.PaymentRepository
}

func NewPaymentService(repo *repositories.PaymentRepository) *PaymentService {
	return &PaymentService{PaymentRepo: repo}
}

// Thêm thanh toán
func (s *PaymentService) AddPayment(payment models.Payment) error {
	return s.PaymentRepo.AddPayment(payment)
}

// Sửa thanh toán
func (s *PaymentService) UpdatePayment(payment models.Payment) error {
	return s.PaymentRepo.UpdatePayment(payment)
}

// Xóa thanh toán
func (s *PaymentService) DeletePayment(paymentID int) error {
	return s.PaymentRepo.DeletePayment(paymentID)
}

// Lấy tất cả thanh toán
func (s *PaymentService) GetAllPayments() ([]models.Payment, error) {
	return s.PaymentRepo.GetAllPayments()
}

// Lấy thanh toán theo ID
func (s *PaymentService) GetPaymentByID(paymentID int) (models.Payment, error) {
	return s.PaymentRepo.GetPaymentByID(paymentID)
}
