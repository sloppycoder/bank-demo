# Dashboard microservice 

A mock mobile banking dashboard gRPC microservice written in Go
 
...more descriptions later...

### Quick Start
1. Install go compiler from [offical site](https://golang.org/dl/go)
2. Install protoc and Go plugin. [Instructions here](https://grpc.io/docs/quickstart/go/)
3. build and run the server and test client
```shell script

make
./bin/server

```

### Tested dependencies

Tested on Ubuntu Linux 20.04 LTS with the following components:

| Component  | Version  |
|---|---|
| [minikube](https://kubernetes.io/docs/setup/learning-environment/minikube/)   | 1.8.2  |
| [Kubernetes](https://kubernetes.io)   | 1.16.6  |
| [Istio](https://istio.io)  | 1.15.2  |
| [skaffold](https://skaffold.dev)  | 1.8.0  |
| [kustomize](https://kustomize.io/)  | 3.5.4  |
| [Go](http://golang.org) compiler | 1.14.2 |
| [protobuf compiler](https://developers.google.com/protocol-buffers) | 3.11.4 |
