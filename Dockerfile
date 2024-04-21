FROM golang:1.21.9-alpine3.19 as builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o gateway cmd/main.go

FROM gcr.io/distroless/static-debian12

WORKDIR /usr/bin

COPY --from=builder /app/gateway .

ENTRYPOINT ["/usr/bin/gateway"]