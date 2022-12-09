package category

import (
	"fp4/models"
	// _product "fp4/service/product"
)

type CategoryResponse struct {
	ID   int64  `json:"id"`
	Type string `json:"type"`
	Spa  int64  `json:"spa"`
	// Product _product.ProductResponse `json:"product,omitempty"`
}

func NewCategoryResponse(category models.Category) CategoryResponse {
	return CategoryResponse{
		ID:   category.ID,
		Type: category.Type,
		Spa:  category.Spa,
	}
}

func CategoryResponseByID(category models.Category) CategoryResponse {
	return CategoryResponse{
		ID:   category.ID,
		Type: category.Type,
		Spa:  category.Spa,
		// Product: _product.NewProductResponseToCategory(category.Products),
	}
}

func NewCatgoeyArrayResponse(categories []models.Category) []CategoryResponse {
	categoryRes := []CategoryResponse{}
	for _, v := range categories {
		p := CategoryResponse{
			ID:   v.ID,
			Type: v.Type,
			Spa:  v.Spa,
		}
		categoryRes = append(categoryRes, p)
	}
	return categoryRes
}
