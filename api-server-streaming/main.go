package main

import (
	"apiServerStream/src/pb/department"
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc"
)

type Server struct {
	department.DepartmentServiceServer
}

func (s *Server) ListDepartmentMember(req *department.DeparmentMembersRequest, srv department.DepartmentService_ListDepartmentMemberServer) error {
	file, err := os.Open("./data.csv")
	if err != nil {
		return fmt.Errorf("erro on open file. error: %v\n", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ";")
		id, _ := strconv.Atoi(data[0])
		name := data[1]
		email := data[2]
		income, _ := strconv.ParseFloat(data[3], 32)
		departmentId, _ := strconv.Atoi(data[4])
		if int32(departmentId) == req.GetDepartmentId() {
			time.Sleep(time.Second)
			if err := srv.Send(&department.DepartmentMembersResponse{
				Id:           int32(id),
				Name:         name,
				Email:        email,
				Income:       float64(income),
				DepartmentId: int32(departmentId),
			}); err != nil {
				return fmt.Errorf("erro on send data[stream]. error: %v\n", err)
			}
		}
	}
	return nil
}
func main() {
	fmt.Println("starting gRPC server on port :9090")
	listenner, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("error on listen server. error: %v\n", err)
	}
	srv := grpc.NewServer()
	department.RegisterDepartmentServiceServer(srv, &Server{})

	if err := srv.Serve(listenner); err != nil {
		log.Fatalf("error on serve. error: %v", err)
	}
}
