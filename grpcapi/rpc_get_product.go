package grpcapi

import (
	"context"
	"strconv"
	"strings"

	"github.com/souzamoab/products-api-golang/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (productServer *ProductServer) GetProduct(context context.Context, req *pb.ProductRequest) (*pb.ProductResponse, error) {
	productIdString := req.GetId()
	productIdInt, _ := strconv.Atoi(productIdString)

	product, err := productServer.productService.FindProductById(productIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			return nil, status.Errorf(codes.NotFound, err.Error())

		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.ProductResponse{
		Product: &pb.Product{
			ID:          int32(product.ID),
			Title:       product.Title,
			Description: product.Description,
			Price:       product.Price,
			Quantity:    int32(product.Quantity),
		},
	}
	return res, nil
}
