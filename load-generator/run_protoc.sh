#!/bin/sh

python -m grpc_tools.protoc -I ../protos -I . --python_out=. --grpc_python_out=. ../protos/demo-bank.proto

