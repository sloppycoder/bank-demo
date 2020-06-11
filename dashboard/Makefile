build:
	go build -o bin/server cmd/server.go

protoc:
	rm -f api/*.pb.go
	protoc -I ../protos ../protos/demo-bank.proto --go_out=plugins=grpc:api

test: 
	go test -coverprofile=coverage.out

all: protoc build test

clean:
	rm -f bin/server
	rm -f coverage.out
