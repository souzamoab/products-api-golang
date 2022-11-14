package main

import (
	"log"

	"github.com/joho/godotenv"
	client "github.com/souzamoab/products-api-golang/clients"
	"github.com/souzamoab/products-api-golang/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "0.0.0.0:8080"
)

func main() {
	godotenv.Load()
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())

	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	defer conn.Close()

	if false {
		listProductsClient := client.NewListProductsClient(conn)

		args := &pb.GetProductsRequest{}

		listProductsClient.ListProducts(args)
	}

	if false {
		createProductClient := client.NewCreateProductClient(conn)

		args := &pb.CreateProductRequest{
			Title:       "Guitar",
			Description: "Les Paul",
			Price:       1500,
			Quantity:    2,
		}

		createProductClient.CreateProduct(args)
	}

	if false {
		updateProductClient := client.NewUpdateProductClient(conn)

		title := "Slash Guitar"
		args := &pb.UpdateProductRequest{
			ID:    1,
			Title: title,
		}

		updateProductClient.UpdateProduct(args)
	}

	if false {
		getProductClient := client.NewGetProductClient(conn)

		args := &pb.ProductRequest{
			Id: "1",
		}

		getProductClient.GetProduct(args)
	}

	if false {
		deleteProductClient := client.NewDeleteProductClient(conn)

		args := &pb.ProductRequest{
			Id: "1",
		}

		deleteProductClient.DeleteProduct(args)
	}
}
