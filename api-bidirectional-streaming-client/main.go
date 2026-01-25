package main

import (
	"apiBidirecionalStreamingClient/src/pb/shoppingcart"
	"context"
	"fmt"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	connection, err := grpc.NewClient("localhost:9090", grpc.WithTransportCredentials(
		insecure.NewCredentials(),
	))
	if err != nil {
		log.Fatalf("erro on cart client connection. error: %v", err)
	}
	defer connection.Close()
	client := shoppingcart.NewShoppingCartServiceClient(connection)
	stream, err := client.AddItem(context.Background())
	if err != nil {
		log.Fatalf("erro on get channel to strean. error: %v\n", err)
	}
	var carts = []shoppingcart.CreateCartRequest{
		{
			ProductId: 2,
			Quantity:  3,
			UnitPrice: 30.0,
		},
		{
			ProductId: 22,
			Quantity:  1,
			UnitPrice: 20.0,
		},
	}
	watch := make(chan struct{})
	go func() {
		for {
			response, err := stream.Recv()
			if err == io.EOF {
				close(watch)
				return
			}
			if err != nil {
				log.Fatalf("error on receive. error: %v\n", err)
			}
			fmt.Printf("Response: %+v\n", response)
		}
	}()

	for _, cart := range carts {
		if err := stream.Send(&cart); err != nil {
			log.Fatalf("error on send data. error: %v\n", err)
		}
		fmt.Printf("Sended: %+v\n", cart)
	}
	stream.CloseSend()
	<-watch
}
