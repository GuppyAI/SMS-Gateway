# SMS-Gateway

The GuppyAI SMS-Gateway is used for sending and receiving SMS messages in the GuppyAI application.
It will push received messages to a message queue (Azure Service Bus) and pull pending responses from it.

# Can I use this?

Yes, you're free to do with it whatever the GPL-3.0 License permits you to do. 
However, if your question is whether you can use this for any actual use case you have, the answer is probably not (without some major adjustment).

This application was written for a very tight use case. We use it with a Brovi/Huawei 4G USB Dongle E3372-325.
To activate modem mode we used this very helpful guide by Pavel Piatruk (@ezbik): [Huawei E3372-325 'BROVI' and Linux (Ubuntu) - Stick mode.](https://blog.tanatos.org/posts/huawei_e3372h-325_brovi_with_linux_stickmode/)

For using this application your SIM card has to be configured to not require a PIN.

If you happen to have the exact same model of USB Dongle on hand, you can use the `make setup` command to set up the stick
for modem mode before starting the application.

# Configuration

| Environment variable                          | Description                                                      | Default   | Possible values                                                                                                                                   |
|-----------------------------------------------|------------------------------------------------------------------|-----------|---------------------------------------------------------------------------------------------------------------------------------------------------|
| GATEWAY_LOGGING_LEVEL                         | Logging level                                                    | WARN      | TRACE, DEBUG, INFO, WARN, DEBUG                                                                                                                   |
| GATEWAY_MESSAGING_ALLOWLIST                   | List of addresses that are allowed to use the gateway's services | NOT SET   | Comma-separated list of addresses, e.g. "sms://+4900000000,sms://+4911111111"                                                                     |
| GATEWAY_SMS_TRACING                           | Tracing of modem commands                                        | 0 (false) | 0 (false), 1 (true)                                                                                                                               |
| GATEWAY_SMS_MODEM_DEVICE                      | Modem device                                                     | NOT SET   | e.g. /dev/ttyUSB1                                                                                                                                 |
| GATEWAY_SMS_MODEM_BAUD                        | Modem baud                                                       | 115200    | Depends on your hardware                                                                                                                          |
| GATEWAY_SMS_POLLING                           | Time between polls for new SMS messages                          | 5s        | Any duration                                                                                                                                      |
| GATEWAY_SERVICEBUS_RECEIVER_QUEUE             | Queue to receive SMS messages from                               | NOT SET   | A valid queue name                                                                                                                                |
| GATEWAY_SERVICEBUS_SENDER_QUEUE               | Queue to send SMS messages received through the modem to         | NOT SET   | A valid queue name                                                                                                                                |
| GATEWAY_SERVICEBUS_RECEIVER_CONNECTIONSTRING  | Connection String used to connect to the receiver queue          | NOT SET   | A valid azure service bus connection string (Format: "Endpoint=sb://some_bus.example.org/;SharedAccessKeyName=Gateway;SharedAccessKey=SecretKey") |
| GATEWAY_SERVICEBUS_SENDER_CONNECTIONSTRING    | Connection String used to connect to the sender queue            | NOT SET   | A valid azure service bus connection string (Format: "Endpoint=sb://some_bus.example.org/;SharedAccessKeyName=Gateway;SharedAccessKey=SecretKey") |

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

*General prerequisites:*

- Accessible modem device
  - This one is a bit tricky, and you'll probably need to experiment with udev rules and kernel modules to get your modem working correctly. Please look up documentation or tutorials on your modem chip for this. The important thing is that you'll need a modem that is accessibly using a tty device (e.g. /dev/ttyUSB1)
- `dialout` group must be set for your current user, if you're not root
  - `usermod -aG dialout $USER`

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
- Make
- Docker (if you really want to, however docker is not recommended and unsupported)

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
    -e GATEWAY_MESSAGING_ALLOWLIST=sms://<PHONE_NUMBER> \
    -e GATEWAY_SERVICEBUS_RECEIVER_QUEUE=<RECEIVER_QUEUE> \
    -e GATEWAY_SERVICEBUS_SENDER_QUEUE=<SENDER_QUEUE> \
    -e GATEWAY_SERVICEBUS_RECEIVER_CONNECTIONSTRING=<RECEIVER_CONNECTIONSTRING> \
    -e GATEWAY_SERVICEBUS_SENDER_CONNECTIONSTRING=<SENDER_CONNECTIONSTRING> \
    ghcr.io/guppyai/sms-gateway:latest
```

*Known issues:*

- When using rootless containers, only the crun runtime is supported as runc does not support sharing supplementary groups (i.e. the dialout group needed to access the modem)
- The `dialout` group needs to be set to run the container. Even if you are running the container rootful.

_Other container tools like docker or containerd are unsupported!_

### Tags

The CI/CD pipeline in this repository will build container images for every push made.

They will be of the following format: YYMMddHHmmss<BRANCH_NAME>
The manifests stored at these tags will provide an image for each of the following platforms:

- linux/arm64v8
- linux/amd64

Also, the latest built image will have the `latest` tag.
However, beware that the `latest` tag will also be applied to images built on other branches than main. 
Therefore, it is not adviceable to use it for any other usecase than caching.

# Copyright

GuppyAI SMS-Gateway (c) 2024 Lucca Greschner and contributors

SPDX-License-Identifier: GPL-3.0
