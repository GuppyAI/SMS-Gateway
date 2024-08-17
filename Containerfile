FROM golang:1.21.9-alpine3.19 as builder

WORKDIR /app

RUN apk add make

COPY Makefile .
COPY go.mod .
COPY go.sum .

RUN make deps

COPY cmd/ ./cmd
COPY internal/ ./internal

RUN make build

FROM alpine:3.20

USER 1000:1000

WORKDIR /usr/bin

COPY --from=builder --chown=1000:1000 --chmod=100 /app/build/gateway-latest ./gateway

ENTRYPOINT ["/usr/bin/gateway"]