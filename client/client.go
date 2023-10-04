package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"grpc.example.com/proto"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewCalculatorClient(conn)

	// Example 1: Addition
	addRequest := &proto.AddRequest{Num1: 10, Num2: 20}
	addResponse, err := client.Add(context.Background(), addRequest)
	if err != nil {
		log.Fatalf("Error calling Add: %v", err)
	}
	fmt.Printf("Addition Result: %.2f\n", addResponse.Result)

	// Example 2: Subtraction
	subtractRequest := &proto.SubtractRequest{Num1: 20, Num2: 10}
	subtractResponse, err := client.Subtract(context.Background(), subtractRequest)
	if err != nil {
		log.Fatalf("Error calling Subtract: %v", err)
	}
	fmt.Printf("Subtraction Result: %.2f\n", subtractResponse.Result)

	// Example 3: Multiplication
	multiplyRequest := &proto.MultiplyRequest{Num1: 10, Num2: 20}
	multiplyResponse, err := client.Multiply(context.Background(), multiplyRequest)
	if err != nil {
		log.Fatalf("Error calling Multiply: %v", err)
	}
	fmt.Printf("Multiplication Result: %.2f\n", multiplyResponse.Result)

	// Example 4: Division
	divideRequest := &proto.DivideRequest{Num1: 20, Num2: 10}
	divideResponse, err := client.Divide(context.Background(), divideRequest)
	if err != nil {
		log.Fatalf("Error calling Divide: %v", err)
	}
	fmt.Printf("Division Result: %.2f\n", divideResponse.Result)
}
