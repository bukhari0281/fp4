package dto

type RegisterRequest struct {
	Full_name string `json:"full_name" form:"full_name" binding:"required,min=1"`
	Email     string `json:"email" form:"email" binding:"required"`
	Password  string `json:"password" form:"password" binding:"required,min=6"`
	Role      string `json:"role" form:"role" binding:"required"`
	Balance   int64  `json:"balance" form:"balance" binding:"required,min=1"`
}
