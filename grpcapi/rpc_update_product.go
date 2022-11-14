package grpcapi

import (
	"context"
	"strings"

	"github.com/souzamoab/products-api-golang/controllers/dto"
	"github.com/souzamoab/products-api-golang/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (productServer *ProductServer) UpdateProduct(context context.Context, req *pb.UpdateProductRequest) (*pb.ProductResponse, error) {
	productId := req.GetID()
	productDto := &dto.ProductUpdateDTO{
		Title:       req.GetTitle(),
		Description: req.GetDescription(),
		Price:       req.GetPrice(),
		Quantity:    uint(req.GetQuantity()),
	}

	updatedProduct, err := productServer.productService.UpdateProduct(int(productId), productDto)

	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			return nil, status.Errorf(codes.NotFound, err.Error())
		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.ProductResponse{
		Product: &pb.Product{
			ID:          int32(updatedProduct.ID),
			Title:       updatedProduct.Title,
			Description: updatedProduct.Description,
			Price:       updatedProduct.Price,
			Quantity:    int32(updatedProduct.Quantity),
		},
	}

	return res, nil

}
