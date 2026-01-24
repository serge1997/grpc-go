package main

import (
	"apiClients/src/pb/products"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	connection, err := grpc.NewClient("localhost:9090", grpc.
		WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalln("Failed to connect to gRPC server[Product]:", err)
	}
	defer connection.Close()
	productClient := products.NewProductServiceClient(connection)
	//Create(productClient, product)

	productList, err := productClient.FindAll(context.Background(), &products.Product{})
	if err != nil {
		log.Fatalln("Failed to retrieve products:", err)
	}

	fmt.Printf("Product List: %+v\n", productList.Products)
}

func Create(client products.ProductServiceClient) {
	product := &products.Product{
		Name:        "Iphone 14",
		Description: "This is am iphone 14, [Create from gRPC client]",
		Price:       19.99,
	}
	productResponse, err := client.Create(context.Background(), product)
	if err != nil {
		log.Fatalln("Failed to create product:", err)
	}

	fmt.Printf("Product Created: %+v\n", productResponse)
}
