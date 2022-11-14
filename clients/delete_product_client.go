package client

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/souzamoab/products-api-golang/pb"
	"google.golang.org/grpc"
)

type DeleteProductClient struct {
	service pb.ProductServiceClient
}

func NewDeleteProductClient(conn *grpc.ClientConn) *DeleteProductClient {
	service := pb.NewProductServiceClient(conn)

	return &DeleteProductClient{service}
}

func (deleteProductClient *DeleteProductClient) DeleteProduct(args *pb.ProductRequest) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*5000))
	defer cancel()

	_, err := deleteProductClient.service.DeleteProduct(ctx, args)

	if err != nil {
		log.Fatalf("DeleteProduct: %v", err)
	}

	fmt.Println("Product deleted successfully")
}
