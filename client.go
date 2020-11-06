package grpcx

import (
	"google.golang.org/grpc"
)

func Dial() (*grpc.ClientConn, error) {

	opts := make([]grpc.DialOption, 0)
	if grpcConfig.Client.Insecure {
		opts = append(opts, grpc.WithInsecure())
	}
	if grpcConfig.Client.Block {
		opts = append(opts, grpc.WithBlock())
	}
	// Set up a connection to the server.
	conn, err := grpc.Dial(ClientConfig().Addr, opts...)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func Close(conn *grpc.ClientConn) {
	if conn != nil {
		if err := conn.Close(); err != nil {
			log.Errorf("close grpc connection failed, err=%v", err)
		}
	}
}
