package messaging

import "errors"

var UnsupportedSchemaErr = errors.New("unsupported schema")

type MessageChannel interface {
	GetSupportedSchema() AddressSchema
	SendMessage(Message) error
	ReceiveMessages(Broker) error
}
