package client

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/souzamoab/products-api-golang/pb"
	"google.golang.org/grpc"
)

type GetProductClient struct {
	service pb.ProductServiceClient
}

func NewGetProductClient(conn *grpc.ClientConn) *GetProductClient {
	service := pb.NewProductServiceClient(conn)

	return &GetProductClient{service}
}

func (getProductClient *GetProductClient) GetProduct(args *pb.ProductRequest) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*5000))
	defer cancel()

	res, err := getProductClient.service.GetProduct(ctx, args)

	if err != nil {
		log.Fatalf("GetProduct: %v", err)
	}

	fmt.Println(res)
}
