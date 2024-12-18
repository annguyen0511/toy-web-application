package services

import (
	"backend/models"
	"backend/repositories"
)

type CartService struct {
	CartRepo *repositories.CartRepository
}

func NewCartService(repo *repositories.CartRepository) *CartService {
	return &CartService{CartRepo: repo}
}

// Thêm giỏ hàng
func (s *CartService) AddCart(cart models.Cart) error {
	return s.CartRepo.AddCart(cart)
}

// Thêm chi tiết giỏ hàng
func (s *CartService) AddCartDetail(cartDetail models.CartDetail) error {
	return s.CartRepo.AddCartDetail(cartDetail)
}

// Sửa chi tiết giỏ hàng
func (s *CartService) UpdateCartDetail(cartDetail models.CartDetail) error {
	return s.CartRepo.UpdateCartDetail(cartDetail)
}

// Xóa chi tiết giỏ hàng
func (s *CartService) DeleteCartDetail(cartDetailID int) error {
	return s.CartRepo.DeleteCartDetail(cartDetailID)
}

// Lấy tất cả chi tiết giỏ hàng theo UserID
func (s *CartService) GetCartByUserID(userID int) ([]models.CartDetail, error) {
	return s.CartRepo.GetCartByUserID(userID)
}

// Lấy tất cả giỏ hàng
func (s *CartService) GetAllCarts() ([]models.Cart, error) {
	return s.CartRepo.GetAllCarts()
}

// Xóa giỏ hàng
func (s *CartService) DeleteCart(cartID int) error {
	return s.CartRepo.DeleteCart(cartID)
}
