package grpcapi

import (
	"context"
	"strings"

	"github.com/souzamoab/products-api-golang/controllers/dto"
	"github.com/souzamoab/products-api-golang/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (productServer *ProductServer) CreateProduct(context context.Context, req *pb.CreateProductRequest) (*pb.ProductResponse, error) {
	productDto := &dto.ProductCreateDTO{
		Title:       req.GetTitle(),
		Description: req.GetDescription(),
		Price:       req.GetPrice(),
		Quantity:    uint(req.GetQuantity()),
	}

	newProduct, err := productServer.productService.CreateProduct(productDto)

	if err != nil {
		if strings.Contains(err.Error(), "title already exists") {
			return nil, status.Errorf(codes.AlreadyExists, err.Error())
		}

		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.ProductResponse{
		Product: &pb.Product{
			ID:          int32(newProduct.ID),
			Title:       newProduct.Title,
			Description: newProduct.Description,
			Price:       newProduct.Price,
			Quantity:    int32(newProduct.Quantity),
		},
	}

	return res, nil

}
