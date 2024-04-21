# SMS Gateway

# How to build and run

## Native

```shell
go mod download
go build -o ./out/gateway ./cmd/main.go
```

## Container

*Building:*

```shell
buildah bud -t gateway:latest # OR
podman build . -t gateway:latest # OR
docker build . -t gateway:latest
```

*Running:*

```shell
podman run --group-add keep-groups \
    --security-opt label=disable \
    --device=/dev/ttyUSB1 \
    -e GATEWAY_SMS_MODEM_BAUD=115200 \
    -e GATEWAY_SMS_MODEM_DEVICE=/dev/ttyUSBxy \
    -e GATEWAY_LOGGING_LEVEL=info \
    -e GATEWAY_MESSAGING_ALLOWLIST=sms://<PHONE_NUMBER> \
    gateway:latest
```