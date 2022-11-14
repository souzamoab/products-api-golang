package services

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/souzamoab/products-api-golang/controllers/dto"
	"github.com/souzamoab/products-api-golang/database"
	"github.com/souzamoab/products-api-golang/entity"
	"gorm.io/gorm"
)

type ProductServiceImpl struct {
	db      *gorm.DB
	context *gin.Context
}

func NewProductService(db *gorm.DB, context *gin.Context) ProductService {
	return &ProductServiceImpl{db, context}
}

func (p *ProductServiceImpl) CreateProduct(productCreateDTO *dto.ProductCreateDTO) (*entity.Product, error) {
	db := database.GetDatabase()

	if err := p.context.ShouldBindJSON(&productCreateDTO); err != nil {
		return nil, err
	}

	product := entity.Product{
		Title:       productCreateDTO.Title,
		Description: productCreateDTO.Description,
		Price:       productCreateDTO.Price,
		Quantity:    productCreateDTO.Quantity}

	if err := db.Create(&product).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *ProductServiceImpl) UpdateProduct(id int, productUpdateDTO *dto.ProductUpdateDTO) (*entity.Product, error) {
	db := database.GetDatabase()

	var product entity.Product

	if err := db.Where("id = ?", id).First(&product).Error; err != nil {
		return nil, err
	}

	if err := p.context.ShouldBindJSON(&productUpdateDTO); err != nil {
		return nil, err
	}

	product.Title = productUpdateDTO.Title
	product.Description = productUpdateDTO.Description
	product.Price = productUpdateDTO.Price
	product.Quantity = productUpdateDTO.Quantity

	db.Save(&product)

	return &product, nil
}

func (p *ProductServiceImpl) DeleteProduct(id int) error {

	db := database.GetDatabase()

	var product entity.Product

	err := db.Where("id = ?", id).First(&product).Error
	if err != nil {
		return gorm.ErrInvalidDB
	}

	err = db.Delete(&product).Error
	if err != nil {
		return errors.New("Not found")
	}

	return nil
}

func (p *ProductServiceImpl) FindProductById(id int) (*entity.Product, error) {
	db := database.GetDatabase()

	var product entity.Product

	err := db.Where("id = ?", id).First(&product).Error
	if err != nil {
		return nil, err
	}

	return &product, err
}

func (p *ProductServiceImpl) FindProducts(page int, limit int) ([]*entity.Product, error) {
	db := database.GetDatabase()

	var products []*entity.Product
	err := db.Find(&products).Error

	if err != nil {
		return nil, err
	}

	return products, err
}
