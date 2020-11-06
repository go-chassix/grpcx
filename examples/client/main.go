package main

import (
	"context"
	"log"
	"os"

	"c5x.io/chassix"
	"c5x.io/grpcx"
	pb "c5x.io/grpcx/examples/helloword"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	chassix.Init()
	log.Printf("cient config: %+v", grpcx.ClientConfig())
	// Set up a connection to the server.
	conn, err := grpcx.Dial()
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer grpcx.Close(conn)
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	log.Printf("timeout: %s", grpcx.ClientConfig().Timeout)
	ctx, cancel := context.WithTimeout(context.Background(), grpcx.ClientConfig().Timeout)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
