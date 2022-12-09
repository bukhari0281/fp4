package dto

type CreateTransactionRequest struct {
	ProductID uint64 `json:"product_id" form:"product_id" binding:"required"`
}
