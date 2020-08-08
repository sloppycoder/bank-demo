FROM python:3-slim-buster as base

FROM base as builder
RUN apt-get -qq update \
    && apt-get install -y --no-install-recommends \
        file \
        g++ \
        libffi-dev

COPY requirements.txt .
RUN pip install --root="/install" -r requirements.txt

FROM golang:1.14 as builder2
WORKDIR /build
ADD . /build
RUN go build -o slave slave.go

FROM base
COPY --from=builder /install /
COPY --from=builder2 /build/slave /

COPY entrypoint.sh *.py /

EXPOSE 8089
ENTRYPOINT ./entrypoint.sh

