package dto

type CreateCategoryRequest struct {
	Type string `gorm:"type" form:"type" binding:"required"`
	Spa  int64  `gorm:"spa" form:"spa"`
}

type UpdateCategoryRequest struct {
	ID                  int64  `gorm:"id" form:"id"`
	Type                string `gorm:"type" form:"type" binding:"required"`
	Sold_product_amount int64  `gorm:"sold_product_amount" form:"sold_product_amount"`
}
