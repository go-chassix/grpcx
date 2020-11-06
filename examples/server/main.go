package main

import (
	"context"
	"log"

	pb "c5x.io/grpcx/examples/helloword"
	"google.golang.org/grpc"

	"c5x.io/chassix"
	"c5x.io/grpcx"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	chassix.Init()
	grpcx.Serve(registerServer)
}

func registerServer(s *grpc.Server) {
	pb.RegisterGreeterServer(s, &server{})
}
