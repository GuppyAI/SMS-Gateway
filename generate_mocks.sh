#!/usr/bin/env bash

go install go.uber.org/mock/mockgen@latest

# gsm package

mockgen -source=./internal/gsm/modem.go -package=gsm -destination=./internal/gsm/mock_modem.go
mockgen -source=./internal/gsm/modem_provider.go -package=gsm -destination=./internal/gsm/mock_modem_provider.go

# sms package

mockgen -source=./internal/sms/receiver.go -package=sms -destination=./internal/sms/mock_receiver.go
mockgen -source=./internal/sms/sender.go -package=sms -destination=./internal/sms/mock_sender.go

# messaging package

mockgen -source=./internal/messaging/broker.go -package=messaging -destination=./internal/messaging/mock_broker.go