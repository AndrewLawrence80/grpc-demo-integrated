package greeter

import (
	"context"
	"fmt"

	pb "github.com/andrewlawrence80/grpc-demo-integrated/proto/greeter"
)

type GreeterServer struct {
	pb.UnimplementedGreeterServer
}

func (s GreeterServer) SayHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Greeting: fmt.Sprintf("Hello %s", request.Name),
	}, nil
}
