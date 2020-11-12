protoc:
	@protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    helloword/helloword.proto
protoc1:
	@protoc --go_out=plugins=grpc:. --go_opt=paths=source_relative helloword/helloword.proto