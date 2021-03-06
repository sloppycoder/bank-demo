## TODOs:
1. enable mTLS between services and setup rules for proper access control (isto 1.5 and 1.4 has different APIs). done (swtiched to 1.4.8 on minikube)
2. add logic to correlate log to trace entries, both for stackdriver and jaeger/elastic (done. added trace and span id to log output in dashboard and casa-account-v1. stackdriver correlation logic is not yet implemetned)
3. oauth token for gRPC service call
4. setup istio authentication rules to examine oauth header for gRPC call
5. expose metrics from microservice and send to both stackdriber and prometheus. (done in dashboard. customer-v1 has some conflict between spring actuator and opencensus, need investigation)
6. log gRPC call service paremeter using opencensus. (possible in Python, no simple solution exist for java and go)
7. write new customer service (done, spring boot)
8. configure egress-gateway for outgoing configuration to database (does not work in minikube)
9. upgrade testdata module's python cassandra driver to solve tls issue on Ubuntu Linux (done)

## Nice to haves
1. add health_probe to casa-account-v1
2. exclude health_probe from trace output
3. enable HTTPS on istio gateway
4. create REST to gRPC gateway (use [envoy gRPC-JSON transcoder](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/grpc_json_transcoder_filter) or 3rd party product like [Ambassador gateway](https://www.getambassador.io/) )
5. create github actions to auto build container images on push to develop
6. handle deployment using Google Application Manager
7. migrate casa-account-v2 to google distroless base image
8. restore the opencensus exporter stackdriver code for casa-account-v2


