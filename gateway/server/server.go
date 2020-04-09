package server

import (
	"context"
	"fmt"
	"log"
	"net"

	addsub "github.com/piotrkira/microservices-calc/addsub/client"
	"github.com/piotrkira/microservices-calc/gateway/endpoints"
	muldiv "github.com/piotrkira/microservices-calc/muldiv/client"

	"google.golang.org/grpc"
)

type Server struct {
	addsubClient *addsub.Client
	muldivClient *muldiv.Client
}

func Run() {
	log.Println("gateway api started")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 7777))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := &Server{addsubClient: addsub.New("addsub"), muldivClient: muldiv.New("muldiv")}
	grpcServer := grpc.NewServer()
	endpoints.RegisterGatewayServer(grpcServer, server)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to listen: %s", err)
	}
}

func (s *Server) Add(ctx context.Context, n *endpoints.Numbers) (*endpoints.Result, error) {
	log.Println("addition requested")
	content, err := s.addsubClient.Add(n.A, n.B)
	if err != nil {
		return &endpoints.Result{Content: ""}, err
	}
	return &endpoints.Result{Content: content}, nil
}

func (s *Server) Sub(ctx context.Context, n *endpoints.Numbers) (*endpoints.Result, error) {
	log.Println("subtraction requested")
	content, err := s.addsubClient.Sub(n.A, n.B)
	if err != nil {
		return &endpoints.Result{Content: ""}, err
	}
	return &endpoints.Result{Content: content}, nil
}

func (s *Server) Mul(ctx context.Context, n *endpoints.Numbers) (*endpoints.Result, error) {
	log.Println("multiplication requested")
	content, err := s.muldivClient.Mul(n.A, n.B)
	if err != nil {
		return &endpoints.Result{Content: ""}, err
	}
	return &endpoints.Result{Content: content}, nil
}

func (s *Server) Div(ctx context.Context, n *endpoints.Numbers) (*endpoints.Result, error) {
	log.Println("division requested")
	content, err := s.muldivClient.Div(n.A, n.B)
	if err != nil {
		return &endpoints.Result{Content: ""}, err
	}
	return &endpoints.Result{Content: content}, nil
}
