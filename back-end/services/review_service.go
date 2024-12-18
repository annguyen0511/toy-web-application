package services

import (
	"backend/models"
	"backend/repositories"
)

type ReviewService struct {
	ReviewRepo *repositories.ReviewRepository
}

func NewReviewService(repo *repositories.ReviewRepository) *ReviewService {
	return &ReviewService{ReviewRepo: repo}
}

// Thêm đánh giá
func (s *ReviewService) AddReview(review models.Review) error {
	return s.ReviewRepo.AddReview(review)
}

// Sửa đánh giá
func (s *ReviewService) UpdateReview(review models.Review) error {
	return s.ReviewRepo.UpdateReview(review)
}

// Xóa đánh giá
func (s *ReviewService) DeleteReview(reviewID int) error {
	return s.ReviewRepo.DeleteReview(reviewID)
}

// Lấy tất cả đánh giá
func (s *ReviewService) GetAllReviews() ([]models.Review, error) {
	return s.ReviewRepo.GetAllReviews()
}

// Lấy đánh giá theo ID sản phẩm
func (s *ReviewService) GetReviewsByProductID(productID int) ([]models.Review, error) {
	return s.ReviewRepo.GetReviewsByProductID(productID)
}
