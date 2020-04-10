package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/piotrkira/microservices-calc/addsub/endpoints"

	"google.golang.org/grpc"
)

type Server struct{}

func Run() {
	log.Println("addsub service started")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 7777))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := &Server{}
	grpcServer := grpc.NewServer()
	endpoints.RegisterAddSubServer(grpcServer, server)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to listen: %s", err)
	}
}

func (s *Server) Add(ctx context.Context, numbers *endpoints.Numbers) (*endpoints.Result, error) {
	log.Println("addition requested")
	hostName, _ := os.Hostname()
	result := numbers.A + numbers.B
	content := fmt.Sprintf("The result is %d calculated by %v", result, hostName)
	return &endpoints.Result{Content: content}, nil
}

func (s *Server) Sub(ctx context.Context, numbers *endpoints.Numbers) (*endpoints.Result, error) {
	log.Println("subtraction requested")
	hostName, _ := os.Hostname()
	result := numbers.A - numbers.B
	content := fmt.Sprintf("The result is %d calculated by %v", result, hostName)
	return &endpoints.Result{Content: content}, nil
}
