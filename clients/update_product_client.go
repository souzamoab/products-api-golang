package client

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/souzamoab/products-api-golang/pb"
	"google.golang.org/grpc"
)

type UpdateProductClient struct {
	service pb.ProductServiceClient
}

func NewUpdateProductClient(conn *grpc.ClientConn) *UpdateProductClient {
	service := pb.NewProductServiceClient(conn)

	return &UpdateProductClient{service}
}

func (updateProductClient *UpdateProductClient) UpdateProduct(args *pb.UpdateProductRequest) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*5000))
	defer cancel()

	res, err := updateProductClient.service.UpdateProduct(ctx, args)

	if err != nil {
		log.Fatalf("UpdateProduct: %v", err)
	}

	fmt.Println(res)
}
