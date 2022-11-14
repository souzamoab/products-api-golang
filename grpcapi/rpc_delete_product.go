package grpcapi

import (
	"context"
	"strconv"
	"strings"

	"github.com/souzamoab/products-api-golang/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (productServer *ProductServer) DeleteProduct(context context.Context, req *pb.ProductRequest) (*pb.DeleteProductResponse, error) {
	productIdString := req.GetId()
	productIdInt, _ := strconv.Atoi(productIdString)

	if err := productServer.productService.DeleteProduct(productIdInt); err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			return nil, status.Errorf(codes.NotFound, err.Error())
		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.DeleteProductResponse{
		Success: true,
	}

	return res, nil
}
