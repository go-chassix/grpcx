package grpcx

import "time"

type GRPCServerConfig struct {
	Port int
	Addr string
}
type GRPCClientConfig struct {
	Addr     string
	Insecure bool
	Block    bool
	Timeout  time.Duration
}

type GRPCConfig struct {
	Server GRPCServerConfig `yaml:"grpcs"`
	Client GRPCClientConfig `yaml:"grpcc"`
}

var grpcConfig = new(GRPCConfig)

func ClientConfig() GRPCClientConfig {
	return grpcConfig.Client
}
