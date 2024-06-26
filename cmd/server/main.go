package main

import (
	"log"
	"net"

	pbCalculator "github.com/andrewlawrence80/grpc-demo-integrated/api/calculator"
	pbGreeter "github.com/andrewlawrence80/grpc-demo-integrated/api/greeter"
	calculator "github.com/andrewlawrence80/grpc-demo-integrated/internal/calculator"
	greeter "github.com/andrewlawrence80/grpc-demo-integrated/internal/greeter"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	greeterServer := greeter.GreeterServer{}
	calculatorServer := calculator.CalculatorServer{}

	pbGreeter.RegisterGreeterServer(s, greeterServer)
	pbCalculator.RegisterCalculatorServer(s, calculatorServer)

	reflection.Register(s)

	log.Println("Server is running on port :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
