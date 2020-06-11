## Run zipkin server in GKE
the server will collect tracing from microservices and forward them to Google Cloud Tracing.

Currently the microservices uses zipkin tracer for easy integration with spring-cloud-sleuth library. There is poentially room for swtiching to a OpenCensus or OpenTracing based implementation which will eliminated the need for this intermediary setup.

## Use ext-dns to update Cloud DNS automatically whenever a new service is exposed
```
kubectl create secret generic gcloud-config --from-file=gcloud-config.json -n default
kubectl apply -f ext-dns.yaml -n default
```