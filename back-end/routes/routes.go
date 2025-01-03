package routes

import (
	"backend/models"

	"github.com/gin-gonic/gin"
)

func SetupRouter(db *models.DB) *gin.Engine {
	r := gin.Default()

	// Gọi các route cho người dùng
	UserRoutes(r, db)

	// Gọi các route cho danh mục
	CategoryRoutes(r, db)

	// Gọi các route cho sản phẩm
	ProductRoutes(r, db)

	// Gọi các route cho đánh giá review
	ReviewRoutes(r, db)

	// Gọi các route cho cart giỏ hàng
	CartRoutes(r, db)

	// Gọi các route cho đơn hàng
	OrderRoutes(r, db)

	// Gọi các route cho thanh toán
	PaymentRoutes(r, db)
	return r
}
