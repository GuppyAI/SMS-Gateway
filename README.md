# SMS Gateway

# Configuration

| Environment variable        | Description                                                      | Default   | Possible values                                                               |
|-----------------------------|------------------------------------------------------------------|-----------|-------------------------------------------------------------------------------|
| GATEWAY_LOGGING_LEVEL       | Logging level                                                    | WARN      | TRACE, DEBUG, INFO, WARN, DEBUG                                               |
| GATEWAY_MESSAGING_ALLOWLIST | List of addresses that are allowed to use the gateway's services | NOT SET   | Comma-separated list of addresses, e.g. "sms://+4900000000,sms://+4911111111" |
| GATEWAY_SMS_TRACING         | Tracing of modem commands                                        | 0 (false) | 0 (false), 1 (true)                                                           |
| GATEWAY_SMS_MODEM_DEVICE    | Modem device                                                     | NOT SET   | e.g. /dev/ttyUSB1                                                             |
| GATEWAY_SMS_MODEM_BAUD      | Modem baud                                                       | 115200    | Depends on your hardware                                                      |
| GATEWAY_SMS_POLLING         | Time between polls for new SMS messages                          | 5s        | Any duration                                                                  | 

# Testing

*Unit tests:*

```shell
make test
```

*Unit tests with coverage report:*

```shell
make coverage
```

The coverage report can be found in the "coverage" folder. "coverage.txt" contains a human-readable format, while "coverage.html" can be viewed in the browser.

# How to build and run

Building has only been tested on Linux with an amd64 processor.
Other operating systems or processor architectures are currently unsupported even though they'll likely compile the code just fine.

## Native

*Prerequisites:*

- Go 1.21 or newer
- Make

*Building:*

```shell
make build
```

The built binary can be found in the "build" folder.

## Container

*Prerequisites:*

- Buildah (preferred for building)
- Podman 4 or newer (preferred for running)
- Docker (if you really want to)

*Building:*

```shell
make container # Use buildah to build the image
make container/podman # Use podman to build the image
make container/docker # Use docker to build the image
```

*Running:*

```shell
podman run --group-add keep-groups \
    --security-opt label=disable \
    --device=/dev/ttyUSBxy \
    -e GATEWAY_SMS_MODEM_BAUD=115200 \
    -e GATEWAY_SMS_MODEM_DEVICE=/dev/ttyUSBxy \
    -e GATEWAY_LOGGING_LEVEL=info \
    -e GATEWAY_MESSAGING_ALLOWLIST=sms://<PHONE_NUMBER> \
    gateway:latest
```

_Other container runtimes are not supported. However, the docker run command should be pretty similar to this one._

# Copyright

GuppyAI SMS-Gateway (c) 2024 Lucca Greschner and contributors

SPDX-License-Identifier: GPL-3.0