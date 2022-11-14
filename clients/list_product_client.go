package client

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/souzamoab/products-api-golang/pb"
	"google.golang.org/grpc"
)

type ListProductsClient struct {
	service pb.ProductServiceClient
}

func NewListProductsClient(conn *grpc.ClientConn) *ListProductsClient {
	service := pb.NewProductServiceClient(conn)

	return &ListProductsClient{service}
}

func (listProductsClient *ListProductsClient) ListProducts(args *pb.GetProductsRequest) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*5000))
	defer cancel()

	stream, err := listProductsClient.service.GetProducts(ctx, args)
	if err != nil {
		log.Fatalf("ListProducts: %v", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("ListProducts: %v", err)
		}

		fmt.Println(res)
	}

}
