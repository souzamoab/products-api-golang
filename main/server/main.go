package main

import (
	"log"
	"net"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/souzamoab/products-api-golang/database"
	"github.com/souzamoab/products-api-golang/grpcapi"
	"github.com/souzamoab/products-api-golang/pb"
	"github.com/souzamoab/products-api-golang/server/routes"
	"github.com/souzamoab/products-api-golang/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	database.StartDatabase()
	godotenv.Load()
	router := routes.ConfigRoutes()
	router.Run(":8080")
}

func startGrpcServer() {
	db := database.GetDatabase()

	productService := services.NewProductService(db, &gin.Context{})

	productServer, err := grpcapi.NewGrpcProductServer(db, productService)
	if err != nil {
		log.Fatal("cannot create grpc productServer: ", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterProductServiceServer(grpcServer, productServer)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatal("cannot create grpc server: ", err)
	}

	log.Printf("start gRPC server on %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot create grpc server: ", err)
	}
}
