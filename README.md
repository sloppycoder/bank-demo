# A modern microservice based application  

This is a demo banking applicaiton built around modern technolgies. The key pillars of this application are:
* Self-contained [microservices](https://microservices.io/patterns/microservices.html) deployed into [Kubernetes](https://kubernetes.io)
* [Polyglot](https://en.wikipedia.org/wiki/Polyglot_(computing)) applicatoin with microservices written in different languages and communicate with each other over standard network protocols.
* Every microservice supports observability with metrics and distributed tracing. Use [opencensus](https://opencensus.io/) and [Prometheus](https://prometheus.io/) to be compatbile with both on-premise stack with [Jaeger](https://www.jaegertracing.io/), [Prometheus](https://prometheus.io/) and [Elastic](https://www.elastic.co/); or [Stackdriver](https://cloud.google.com/products/operations) in GCP.
* Use [Istio](https://istio.io) service mesh to perform traffic management within and inter kubernetes clusters
* High performance inter microservice communication with [gRPC](https://grpc.io). The gRPC services can be directly exposed to external client or use [Envoy gRPC-Web filter](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/grpc_web_filter) for clients that do not support HTTP/2 natively.



![architecture diagram](doc/architecture.png)





## Modules in this repo

| Directory     | Conent      |
| ------------- |-------------| 
| [protos](protos)    | protobuf API definition used in all across the application. This is the master copy. |
| [dashboard](dashboard) | dashboard microservice, written in Go. It calls casa-account services, either v1 or v2, controlled by [istio virtualservice](https://istio.io/latest/docs/reference/config/networking/virtual-service/) configuration      |
| [casa-account-v1](casa-account-v1) | casa account microservice written in Java using [micronaut] framework. It reads account data stored in Cassandra and return the clients |
| [casa-account-v2](casa-account-v2) | casa account microservice written in Javascript and runs with nodejs. It only returns dummy data. |
| [load-generator](load-generator) | load test script written in Python and uses [Locust](https://locust.io/). The script can be run outside the cluster or deployed into the cluster |
| [testdata](testdata) | Python scripts that generate test data and write them to Cassandra |
| [istio](istio) | istio manifest files to traffice management |
| [gcp](gcp) | script and manifest to provision GKE cluster. |

### Requirement
The code in this repo has been tested with minikube and GKE. Details below

Local Minikue
* Ubuntu Linux 20.04
* [minikube](https://kubernetes.io/docs/setup/learning-environment/minikube/) v1.11.0
* Kubernetes 1.16.9
* Istio 1.5.2
* [Apache Cassandra](https://cassandra.apache.org/) version 4.0.0

Google Kubernetes Engine
 * Kubernetes provisioned using [regular channel provided by GKE](https://cloud.google.com/kubernetes-engine/docs/release-notes-regular), v1.16.8-gke.15 as of now
 * Istio is provisioned as a feature of GKE cluster, 1.4.6-gke.0 as of now
 * [Datastax Astra](https://www.datastax.com/products/datastax-astra) running in GCP

Development Tools:
* [kustomize](https://github.com/kubernetes-sigs/kustomize) 3.5.4
* [skaffold](https://skaffold.dev) v1.10.1

### Roadmap
1. better integration of logging and tracing. ability to navigate from trace to log with high accuracy.
2. use an external OpenID Connect provide for authentication.
3. enable mTLS for inter microservice communication.
4. use Istio for oauth token validation


