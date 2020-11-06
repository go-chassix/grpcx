package grpcx

import (
	"google.golang.org/grpc"
	"net"
	"strconv"

	"c5x.io/chassix"
	"c5x.io/logx"
)

var log *logx.Entry

func init() {
	chassix.Register(&chassix.Module{Name: chassix.ModuleGrpc,
		ConfigPtr: grpcConfig})
	log = logx.New().Category("chassix").Component("grpc")
}

type RegisterServerFunc func(server *grpc.Server)

func Serve(rsFN RegisterServerFunc) {
	addr := grpcConfig.Server.Addr + ":" + strconv.Itoa(grpcConfig.Server.Port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	rsFN(s)
	log.Infof("grpc server start on %s", addr)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
