package grpcapi

import (
	"github.com/souzamoab/products-api-golang/pb"
	"github.com/souzamoab/products-api-golang/services"
	"gorm.io/gorm"
)

type ProductServer struct {
	pb.UnimplementedProductServiceServer
	db             *gorm.DB
	productService services.ProductService
}

func NewGrpcProductServer(db *gorm.DB, productService services.ProductService) (*ProductServer, error) {
	productServer := &ProductServer{
		db:             db,
		productService: productService,
	}

	return productServer, nil
}
