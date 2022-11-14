package grpcapi

import (
	"github.com/souzamoab/products-api-golang/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (productServer *ProductServer) GetProducts(req *pb.GetProductsRequest, stream pb.ProductService_GetProductsServer) error {
	var page = req.GetPage()
	var limit = req.GetLimit()

	products, err := productServer.productService.FindProducts(int(page), int(limit))
	if err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}

	for _, product := range products {
		stream.Send(&pb.Product{
			ID:          int32(product.ID),
			Title:       product.Title,
			Description: product.Description,
			Price:       product.Price,
			Quantity:    int32(product.Quantity),
		})
	}

	return nil
}
