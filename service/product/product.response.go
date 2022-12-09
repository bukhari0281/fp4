package _product

import (
	"fp4/models"
	_category "fp4/service/category"
	_user "fp4/service/user"
)

type ProductResponse struct {
	ID       int64                      `json:"id"`
	Title    string                     `json:"title"`
	Price    int64                      `json:"price"`
	Stoct    int64                      `json:"stoct"`
	Category _category.CategoryResponse `json:"category,omitempty"`
	User     _user.UserResponse         `json:"user,omitempty"`
}

func NewProductResponse(product models.Product) ProductResponse {
	return ProductResponse{
		ID:       product.ID,
		Title:    product.Title,
		Price:    product.Price,
		Stoct:    product.Stoct,
		Category: _category.NewCategoryResponse(product.Category),
		User:     _user.NewUserResponse(product.User),
	}
}

func NewProductResponseToCategory(product models.Product) ProductResponse {
	return ProductResponse{
		ID:    product.ID,
		Title: product.Title,
		Price: product.Price,
		Stoct: product.Stoct,
		// Category: _category.NewCategoryResponse(product.Category),
	}
}

func NewProductArrayResponse(products []models.Product) []ProductResponse {
	productRes := []ProductResponse{}
	for _, v := range products {
		p := ProductResponse{
			ID:       v.ID,
			Title:    v.Title,
			Price:    v.Price,
			Stoct:    v.Stoct,
			Category: _category.CategoryResponse{},
		}
		productRes = append(productRes, p)
	}
	return productRes
}
