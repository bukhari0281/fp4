package repo

import (
	"fp4/models"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	// All(userID string) (*[]_category.CategoryResponse, error)
	InsertCategory(category models.Category) (models.Category, error)
	FindOneCategoryByID(ID string) (models.Category, error)
	// FindAllproductByCategoryID(productID string) ()
}

type categoryRepo struct {
	connection *gorm.DB
}

// FindOneCategoryByID implements CategoryRepository
func (c *categoryRepo) FindOneCategoryByID(categoryID string) (models.Category, error) {
	var category models.Category
	res := c.connection.Preload("User").Where("id = ?", categoryID).Take(&category)
	if res.Error != nil {
		return category, res.Error
	}
	return category, nil
}

func NewCategoryRepo(connection *gorm.DB) CategoryRepository {
	return &categoryRepo{
		connection: connection,
	}
}

// InsertCategory implements CategoryRepository
func (c *categoryRepo) InsertCategory(category models.Category) (models.Category, error) {
	c.connection.Save(&category)
	c.connection.Preload("User").Find(&category)
	return category, nil
}
