package services

import (
	"errors"

	database "github.com/hosnibounechada/go-api/internal/db"
	"github.com/hosnibounechada/go-api/internal/product/models"
	"github.com/jinzhu/gorm"
)

type ProductService struct{}

func NewProductService() *ProductService {
	return &ProductService{}
}

func (s *ProductService) GetAll() ([]models.Product, error) {
	var products []models.Product

	if err := database.DB.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (s *ProductService) Get(id int64) (models.Product, error) {
	existingProduct := models.Product{ID: id}

	result := database.DB.First(&existingProduct)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return models.Product{}, errors.New("Product not found")
		}
		return models.Product{}, result.Error
	}

	return existingProduct, nil
}

func (s *ProductService) Create(product models.CreateProductDTO) (models.Product, error) {
	newProduct := models.Product{
		Name:     product.Name,
		Price:    product.Price,
		Quantity: product.Quantity,
		UserID:   product.UserID,
	}

	result := database.DB.Create(&newProduct)
	if result.Error != nil {
		return models.Product{}, result.Error
	}

	return newProduct, nil
}

func (s *ProductService) Update(id int64, product models.UpdateProductDTO) (models.Product, error) {
	existingProduct := models.Product{ID: id}

	result := database.DB.First(&existingProduct)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return models.Product{}, errors.New("Product not found")
		}
		return models.Product{}, result.Error
	}

	existingProduct.Name = product.Name
	existingProduct.Price = product.Price
	existingProduct.Quantity = product.Quantity

	result = database.DB.Save(&existingProduct)
	if result.Error != nil {
		return models.Product{}, result.Error
	}

	return existingProduct, nil
}

func (s *ProductService) Delete(id int64) error {
	result := database.DB.Where("id = ?", id).Delete(&models.Product{})

	println(result)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
