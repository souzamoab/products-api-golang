package services

import (
	"github.com/souzamoab/products-api-golang/controllers/dto"
	"github.com/souzamoab/products-api-golang/entity"
)

type ProductService interface {
	CreateProduct(*dto.ProductCreateDTO) (*entity.Product, error)
	UpdateProduct(int, *dto.ProductUpdateDTO) (*entity.Product, error)
	FindProductById(int) (*entity.Product, error)
	FindProducts(page int, limit int) ([]*entity.Product, error)
	DeleteProduct(int) error
}
