package greeter

import (
	"context"
	"testing"

	pb "github.com/andrewlawrence80/grpc-demo-integrated/api/greeter"
	server "github.com/andrewlawrence80/grpc-demo-integrated/internal/greeter"
	"github.com/golang/mock/gomock"
)

func TestGreeterServer_SayHello(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockGreeterServer := NewMockGreeterServer(mockCtrl)
	mockGreeterServer.EXPECT().SayHello(gomock.Any(), gomock.Any()).AnyTimes().DoAndReturn(
		func(ctx context.Context, request *pb.HelloRequest) (*pb.HelloResponse, error) {
			s := &server.GreeterServer{}
			return s.SayHello(ctx, request)
		},
	)

	type args struct {
		ctx     context.Context
		request *pb.HelloRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *pb.HelloResponse
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "TestSayHello_Alice",
			args: args{
				ctx:     context.Background(),
				request: &pb.HelloRequest{Name: "Alice"},
			},
			want:    &pb.HelloResponse{Greeting: "Hello Alice"},
			wantErr: false,
		},
		{
			name: "TestSayHello_Bob",
			args: args{
				ctx:     context.Background(),
				request: &pb.HelloRequest{Name: "Bob"},
			},
			want:    &pb.HelloResponse{Greeting: "Hello Bob"},
			wantErr: false,
		},
		{
			name:    "TestSayHello_Empty",
			args:    args{ctx: context.Background(), request: &pb.HelloRequest{Name: ""}},
			want:    &pb.HelloResponse{Greeting: "Hello "},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := mockGreeterServer.SayHello(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("GreeterServer.SayHello() error = %v, wantErr %v", err, tt.wantErr)
			}
			if resp != nil && resp.Greeting != tt.want.Greeting {
				t.Errorf("GreeterServer.SayHello() = %v, want %v", resp.Greeting, tt.want.Greeting)
			}
		})
	}
}
