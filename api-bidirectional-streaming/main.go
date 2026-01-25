package main

import (
	"apiBidirecionalStreaming/src/pb/shoppingcart"
	"fmt"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	shoppingcart.ShoppingCartServiceServer
}

func (s *Server) AddItem(srv shoppingcart.ShoppingCartService_AddItemServer) error {
	var quantity int32 = 0
	var total float64 = 0
	for {
		newItem, err := srv.Recv()
		if err == io.EOF {
			//client connection ended here or streaming ended
			return srv.Send(&shoppingcart.ShoppingCartTotal{
				QuantityItems: quantity,
				TotalPrice:    float64(total),
			})
		}
		if err != nil {
			return fmt.Errorf("erro on receive. error: %v\n", err)
		}
		quantity += newItem.GetQuantity()
		total += float64(newItem.GetQuantity()) * newItem.GetUnitPrice()
		//client still connected - streaming is running
		srv.Send(&shoppingcart.ShoppingCartTotal{
			QuantityItems: quantity,
			TotalPrice:    total,
		})

	}
}

func main() {
	fmt.Println("starting server om port 9090")
	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("erro on listen. error: %v\n", err)
	}
	srv := grpc.NewServer()
	shoppingcart.RegisterShoppingCartServiceServer(srv, &Server{})
	if err := srv.Serve(listener); err != nil {
		log.Fatalf("error on serve. %v\n", err)
	}
}
