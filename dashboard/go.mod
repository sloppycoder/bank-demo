module dashboard

go 1.14

require (
	contrib.go.opencensus.io/exporter/prometheus v0.2.0
	contrib.go.opencensus.io/exporter/stackdriver v0.13.1
	contrib.go.opencensus.io/exporter/zipkin v0.1.1
	github.com/golang/protobuf v1.4.0
	github.com/openzipkin/zipkin-go v0.2.2
	github.com/sirupsen/logrus v1.6.0
	github.com/stretchr/testify v1.4.0
	go.opencensus.io v0.22.2
	golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e
	google.golang.org/genproto v0.0.0-20190911173649-1774047e7e51
	google.golang.org/grpc v1.29.1
	google.golang.org/protobuf v1.23.0
)
