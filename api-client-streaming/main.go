package main

import (
	"apiClientStreaming/src/pb/calc"
	"fmt"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	calc.CalServiceServer
}

func (s *Server) Calc(stream calc.CalService_CalcServer) error {
	var quantity int32 = 0
	var total int32 = 0
	for {
		input, err := stream.Recv()
		if err == io.EOF {
			avg := float64(total) / float64(quantity)
			return stream.SendAndClose(&calc.Output{
				Quantity: quantity,
				Total:    total,
				Average:  avg,
			})
		}
		if err != nil {
			return err
		}
		quantity++
		total += input.GetValue()
		fmt.Printf("input: %+v\n", input)
	}
}
func main() {
	log.Println("start Grpc server...")
	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalln("erro on get listener. erro: ", err)
	}
	s := grpc.NewServer()
	calc.RegisterCalServiceServer(s, &Server{})

	if err := s.Serve(listener); err != nil {
		log.Fatalln("erro on build server. erro: ", err)
	}
}
