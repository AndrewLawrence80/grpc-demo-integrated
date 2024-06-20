package main

import (
	"context"
	"log"
	"time"

	pbGreeter "github.com/andrewlawrence80/grpc-demo-integrated/api/greeter"
	"google.golang.org/grpc"
)

func Greet() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pbGreeter.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := client.SayHello(ctx, &pbGreeter.HelloRequest{Name: "World"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetGreeting())
}
