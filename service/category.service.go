package service

import (
	"fp4/dto"
	"fp4/models"
	"fp4/repo"
	_category "fp4/service/category"
	"log"
	"strconv"

	"github.com/mashingan/smapping"
)

type CategoryService interface {
	// All(userID string) (*[]_category.CategoryResponse, error)
	FindOneCategoryByID(categoryID string) (*_category.CategoryResponse, error)
	CreateCategory(categoryRequest dto.CreateCategoryRequest, userID string) (*_category.CategoryResponse, error)
}

type categoryService struct {
	categoryRepo repo.CategoryRepository
}

func NewCategoryService(categoryRepo repo.CategoryRepository) CategoryService {
	return &categoryService{
		categoryRepo: categoryRepo,
	}
}

func (c *categoryService) FindOneCategoryByID(categoryID string) (*_category.CategoryResponse, error) {
	category, err := c.categoryRepo.FindOneCategoryByID(categoryID)

	if err != nil {
		return nil, err
	}

	res := _category.NewCategoryResponse(category)
	return &res, nil
}

// func (c *categoryService) All(*[]_category.CategoryResponse, error) {
// 	categories, err := c.categoryRepo.All(userID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	categoris := _category.NewCatgoeyArrayResponse
// }

// CreateCategory implements CategoryService
func (c *categoryService) CreateCategory(categoryRequest dto.CreateCategoryRequest, userID string) (*_category.CategoryResponse, error) {
	category := models.Category{}
	err := smapping.FillStruct(&category, smapping.MapFields(&categoryRequest))

	if err != nil {
		log.Fatalf("Failed map %v", err)
		return nil, err
	}

	id, _ := strconv.ParseInt(userID, 0, 64)
	category.UserID = id
	p, err := c.categoryRepo.InsertCategory(category)
	if err != nil {
		return nil, err
	}

	res := _category.NewCategoryResponse(p)
	return &res, nil
}
