FROM golang:1.14 as builder

WORKDIR /build
ADD . /build

RUN go build -o bin/server cmd/server.go

FROM gcr.io/distroless/base

COPY --from=builder /build/bin/server /server
# use the local copy. downloading from github for each build is way too slow
COPY --from=builder /build/bin/grpc_health_probe /grpc_health_probe

EXPOSE 50051

CMD ["/server"]
