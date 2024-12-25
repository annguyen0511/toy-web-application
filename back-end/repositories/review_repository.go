package repositories

import (
	"backend/models"
	"database/sql"
	"time"
)

type ReviewRepository struct {
	DB *sql.DB
}

func NewReviewRepository(db *sql.DB) *ReviewRepository {
	return &ReviewRepository{DB: db}
}

// Thêm đánh giá
func (r *ReviewRepository) AddReview(review models.Review) error {
	_, err := r.DB.Exec("INSERT INTO Reviews (ProductID, UserID, Rating, Comment) VALUES (?, ?, ?, ?)",
		review.ProductID, review.UserID, review.Rating, review.Comment)
	return err
}

// Sửa đánh giá
func (r *ReviewRepository) UpdateReview(review models.Review) error {
	_, err := r.DB.Exec("UPDATE Reviews SET Rating = ?, Comment = ? WHERE ReviewID = ?",
		review.Rating, review.Comment, review.ReviewID)
	return err
}

// Xóa đánh giá
func (r *ReviewRepository) DeleteReview(reviewID int) error {
	_, err := r.DB.Exec("DELETE FROM Reviews WHERE ReviewID = ?", reviewID)
	return err
}

// Lấy tất cả đánh giá
func (r *ReviewRepository) GetAllReviews() ([]models.Review, error) {
	rows, err := r.DB.Query("SELECT ReviewID, ProductID, UserID, Rating, Comment, CreatedAt FROM Reviews")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []models.Review
	for rows.Next() {
		var review models.Review
		var createdAt string // Temporary variable to hold the CreatedAt value
		if err := rows.Scan(&review.ReviewID, &review.ProductID, &review.UserID, &review.Rating, &review.Comment, &createdAt); err != nil {
			return nil, err
		}
		// Convert createdAt string to time.Time
		if parsedTime, err := time.Parse("2006-01-02 15:04:05", createdAt); err == nil {
			review.CreatedAt = parsedTime
		} else {
			println("Error parsing CreatedAt:", err.Error()) // Print error
			return nil, err
		}
		reviews = append(reviews, review)
	}
	return reviews, nil
}

// Lấy đánh giá theo ID sản phẩm
func (r *ReviewRepository) GetReviewsByProductID(productID int) ([]models.Review, error) {
	rows, err := r.DB.Query("SELECT ReviewID, ProductID, UserID, Rating, Comment, CreatedAt FROM Reviews WHERE ProductID = ?", productID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []models.Review
	for rows.Next() {
		var review models.Review
		if err := rows.Scan(&review.ReviewID, &review.ProductID, &review.UserID, &review.Rating, &review.Comment, &review.CreatedAt); err != nil {
			return nil, err
		}
		reviews = append(reviews, review)
	}
	return reviews, nil
}
