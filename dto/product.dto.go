package dto

type CreateProductRequest struct {
	Title      string `grom:"title" form:"title" binding:"required"`
	Price      int64  `grom:"price" form:"price" binding:"required"`
	Stoct      int64  `grom:"stoct" form:"stoct" binding:"required,min=1"`
	CategoryID int64  `json:"category_id" form:"category_id" binding:"required"`
}

type UpdateProductRequest struct {
	ID    int64  `json:"id" form:"id"`
	Title string `grom:"title" form:"title" binding:"required"`
	Price int64  `grom:"price" form:"price" binding:"required"`
	Stoct int64  `grom:"stoct" form:"stoct" binding:"required,min=1"`
}
