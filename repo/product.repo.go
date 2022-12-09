package repo

import (
	"fp4/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	All(userID string) ([]models.Product, error)
	InsertProduct(product models.Product) (models.Product, error)
	FindOneProductByID(ID string) (models.Product, error)
	FindAllProduct(userID string) ([]models.Product, error)
	UpdateProduct(product models.Product) (models.Product, error)
	DeleteProduct(productID string) error
}

type productRepo struct {
	connection *gorm.DB
}

func NewProductRepo(connection *gorm.DB) ProductRepository {
	return &productRepo{
		connection: connection,
	}
}

// All implements ProductRepository
// func (c *productRepo) All(userID string) ([]models.Product, error) {
// 	products := []models.Product{}
// 	c.connection.Preload("User").Where("user_id = ?", userID).Find(&products)
// 	return products, nil

// FindOneproductByID implements ProductRepository
//
//	func (c *productRepo) FindOneProductByID(productID string) (models.Product, error) {
//		var product models.Product
//		res := c.connection.Preload("Category").Preload("User").Where("id = ?", productID).Take(&product)
//		if res.Error != nil {
//			return product, res.Error
//		}
//		return product, nil
//	}

func (c *productRepo) All(userID string) ([]models.Product, error) {
	products := []models.Product{}
	c.connection.Preload("User").Where("user_id = ?", userID).Find(&products)
	return products, nil
}

func (c *productRepo) UpdateProduct(product models.Product) (models.Product, error) {
	c.connection.Save(&product)
	c.connection.Preload("User").Find(&product)
	return product, nil
}

func (c *productRepo) FindOneProductByID(productID string) (models.Product, error) {
	var product models.Product
	res := c.connection.Preload("User").Where("id = ?", productID).Take(&product)
	if res.Error != nil {
		return product, res.Error
	}
	return product, nil
}

// InserProdect implements ProductRepository
func (c *productRepo) InsertProduct(product models.Product) (models.Product, error) {
	c.connection.Save(&product)
	c.connection.Preload("User").Preload("Category").Find(&product)
	return product, nil
}

func (c *productRepo) FindAllProduct(userID string) ([]models.Product, error) {
	products := []models.Product{}
	c.connection.Where("user_id = ?", userID).Find(&products)
	return products, nil
}

func (c *productRepo) DeleteProduct(productID string) error {
	var product models.Product
	res := c.connection.Preload("User").Where("id = ?", productID).Take(&product)
	if res.Error != nil {
		return res.Error
	}
	c.connection.Delete(&product)
	return nil
}
