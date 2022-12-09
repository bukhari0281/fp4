package _user

import "fp4/models"

type UserResponse struct {
	ID int64 `json:"id"`
	// Full_name string `json:"full_name"`
	// Email     string `json:"email"`
	// Role      string `json:"role"`
	// Balance   int64  `json:"balance"`
	Token string `json:"token,omitempty"`
}

type UserLoginResponse struct {
	Token string `json:"token,omitempty"`
}

func NewUserResponse(user models.User) UserResponse {
	return UserResponse{
		ID: user.ID,
		// Full_name: user.Full_name,
		// Email:     user.Email,
		// Role:      user.Role,
		// Balance:   user.Balance,
		Token: "",
	}
}
