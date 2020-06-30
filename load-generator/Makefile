build:
	go build -o slave slave.go

protoc:
	rm -f api/*.pb.go
	protoc -I ../protos ../protos/demo-bank.proto --go_out=plugins=grpc:api

all: protoc build test

clean:
	rm -f slave
