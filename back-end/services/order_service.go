package services

import (
	"backend/models"
	"backend/repositories"
)

type OrderService struct {
	OrderRepo *repositories.OrderRepository
}

func NewOrderService(repo *repositories.OrderRepository) *OrderService {
	return &OrderService{OrderRepo: repo}
}

// Thêm đơn hàng
func (s *OrderService) AddOrder(order models.Order) error {
	return s.OrderRepo.AddOrder(order)
}

// Thêm chi tiết đơn hàng
func (s *OrderService) AddOrderDetail(orderDetail models.OrderDetail) error {
	return s.OrderRepo.AddOrderDetail(orderDetail)
}

// Sửa chi tiết đơn hàng
func (s *OrderService) UpdateOrderDetail(orderDetail models.OrderDetail) error {
	return s.OrderRepo.UpdateOrderDetail(orderDetail)
}

// Xóa chi tiết đơn hàng
func (s *OrderService) DeleteOrderDetail(orderDetailID int) error {
	return s.OrderRepo.DeleteOrderDetail(orderDetailID)
}

// Lấy tất cả đơn hàng theo UserID
func (s *OrderService) GetOrdersByUserID(userID int) ([]models.Order, error) {
	return s.OrderRepo.GetOrdersByUserID(userID)
}

// Lấy tất cả đơn hàng
func (s *OrderService) GetAllOrders() ([]models.Order, error) {
	return s.OrderRepo.GetAllOrders()
}

// Lấy tất cả chi tiết đơn hàng theo OrderID
func (s *OrderService) GetOrderDetailsByOrderID(orderID int) ([]models.OrderDetail, error) {
	return s.OrderRepo.GetOrderDetailsByOrderID(orderID)
}

// Lấy tất cả chi tiết đơn hàng
func (s *OrderService) GetAllOrderDetails() ([]models.OrderDetail, error) {
	return s.OrderRepo.GetAllOrderDetails()
}

// Cập nhật trạng thái đơn hàng
func (s *OrderService) UpdateOrderStatus(orderID int, status string) error {
	return s.OrderRepo.UpdateOrderStatus(orderID, status)
}
