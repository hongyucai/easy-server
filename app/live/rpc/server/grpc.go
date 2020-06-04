package grpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "go-mod/protocol/protobuf/grpc" //引入了生成的go代码
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCalculatorServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type server struct {
	pb.UnimplementedCalculatorServer
}

func (s *server) Add(ctx context.Context, in *pb.RequestData) (*pb.ResponseData, error) {
	fmt.Println("add:")
	return &pb.ResponseData{
		C: in.A + in.B,
	}, nil
}
func (s *server) Sub(ctx context.Context, in *pb.RequestData) (*pb.ResponseData, error) {
	fmt.Println("sub:")
	return &pb.ResponseData{
		C: in.A - in.B,
	}, nil
}
