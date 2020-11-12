module c5x.io/grpcx

go 1.14

require (
	c5x.io/chassix v0.0.0-20201106081434-32e93f693be6
	c5x.io/logx v0.0.0-20200904004942-c2201b5c53ab
	github.com/antonfisher/nested-logrus-formatter v1.3.0 // indirect
	github.com/golang/protobuf v1.4.3
	github.com/sirupsen/logrus v1.7.0 // indirect
	golang.org/x/sys v0.0.0-20201106081118-db71ae66460a // indirect
	google.golang.org/grpc v1.29.1
	google.golang.org/protobuf v1.25.0
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.29.1
