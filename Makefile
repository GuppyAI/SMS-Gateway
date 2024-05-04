.PHONY: all mocks test deps container

GO=go
GOTEST=$(GO) test
GOCOVER=$(GO) tool cover
GOMOD=$(GO) mod
GOBUILD=$(GO) build -ldflags="-s -w"
GORUN=$(GO) run

MOCKGEN=~/go/bin/mockgen

CONTAINERBUILD=buildah bud

VERSION=latest

all: test build container

# Downloads dependencies
deps:
	$(GO) mod download

# Builds
build: deps
	$(GOMOD) download
	$(GOBUILD) -o ./build/gateway-$(VERSION) ./cmd/main.go

container:
	$(CONTAINERBUILD) -t gateway:$(VERSION)

test: deps
	$(GOTEST) -v ./...

coverage: coverage/coverage.txt coverage/coverage.html

coverage/coverage.out: deps mocks
	mkdir -p coverage
	$(GOTEST) -v -coverprofile=coverage/coverage.out.tmp ./...
	grep -v "mock_" coverage/coverage.out.tmp > coverage/coverage.out
	rm coverage/coverage.out.tmp

coverage/coverage.txt: coverage/coverage.out
	mkdir -p coverage
	$(GOCOVER) -func=coverage/coverage.out -o=coverage/coverage.txt

coverage/coverage.html: coverage/coverage.out
	mkdir -p coverage
	$(GOCOVER) -html=coverage/coverage.out -o=coverage/coverage.html

mocks:
	$(GO) install go.uber.org/mock/mockgen@latest

	# gsm package

	$(MOCKGEN) -source=./internal/gsm/modem.go -package=gsm -destination=./internal/gsm/mock_modem.go
	$(MOCKGEN) -source=./internal/gsm/modem_provider.go -package=gsm -destination=./internal/gsm/mock_modem_provider.go

    # sms package

	$(MOCKGEN) -source=./internal/sms/receiver.go -package=sms -destination=./internal/sms/mock_receiver.go
	$(MOCKGEN) -source=./internal/sms/sender.go -package=sms -destination=./internal/sms/mock_sender.go

    # messaging package

	$(MOCKGEN) -source=./internal/messaging/broker.go -package=messaging -destination=./internal/messaging/mock_broker.go
	$(MOCKGEN) -source=./internal/messaging/message_handler.go -package=messaging -destination=./internal/messaging/mock_message_handler.go
	$(MOCKGEN) -source=./internal/messaging/message_channel.go -package=messaging -destination=./internal/messaging/mock_message_channel.go