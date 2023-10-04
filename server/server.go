package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"grpc.example.com/proto"
)

type server struct{}

func (s *server) Add(ctx context.Context, req *proto.AddRequest) (*proto.AddResponse, error) {
	a, b := req.Num1, req.Num2
	result := a + b
	return &proto.AddResponse{Result: result}, nil
}

func (s *server) Multiply(ctx context.Context, req *proto.MultiplyRequest) (*proto.MultiplyResponse, error) {
	a, b := req.Num1, req.Num2
	result := a * b
	return &proto.MultiplyResponse{Result: result}, nil
}

func (s *server) Divide(ctx context.Context, req *proto.DivideRequest) (*proto.DivideResponse, error) {
	a, b := req.Num1, req.Num2
	result := a / b
	return &proto.DivideResponse{Result: result}, nil
}

func (s *server) Subtract(ctx context.Context, req *proto.SubtractRequest) (*proto.SubtractResponse, error) {
	a, b := req.Num1, req.Num2
	result := a - b
	return &proto.SubtractResponse{Result: result}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	fmt.Println("Server starting at port :50051")
	grpcServer := grpc.NewServer()
	proto.RegisterCalculatorServer(grpcServer, &server{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
}
