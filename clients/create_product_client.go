package client

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/souzamoab/products-api-golang/pb"
	"google.golang.org/grpc"
)

type CreateProductClient struct {
	service pb.ProductServiceClient
}

func NewCreateProductClient(conn *grpc.ClientConn) *CreateProductClient {
	service := pb.NewProductServiceClient(conn)

	return &CreateProductClient{service}
}

func (createProductClient *CreateProductClient) CreateProduct(args *pb.CreateProductRequest) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*5000))
	defer cancel()

	res, err := createProductClient.service.CreateProduct(ctx, args)

	if err != nil {
		log.Fatalf("CreateProduct: %v", err)
	}

	fmt.Println(res)
}
