package greeter

import (
	context "context"
	"net"
	"testing"

	pb "github.com/andrewlawrence80/grpc-demo-integrated/api/greeter"
	"github.com/andrewlawrence80/grpc-demo-integrated/internal/greeter"
	"github.com/stretchr/testify/assert"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, greeter.GreeterServer{})
	go func() {
		if err := s.Serve(lis); err != nil {
			panic(err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestGreeterClient_SayHello(t *testing.T) {
	conn, err := grpc.DialContext(context.Background(), "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	assert.NoError(t, err)
	defer conn.Close()

	client := pb.NewGreeterClient(conn)
	req := &pb.HelloRequest{Name: "World"}
	res, err := client.SayHello(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, "Hello World", res.Greeting)
}
