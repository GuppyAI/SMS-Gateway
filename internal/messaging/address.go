package messaging

import (
	"errors"
	"strings"
)

type AddressSchema string

var ErrInvalidAddressFormat = errors.New("invalid address format")

const (
	SMS   AddressSchema = "sms"
	EMail               = "email"
)

type Address struct {
	schema  AddressSchema
	address string
}

func NewAddress(schema AddressSchema, address string) *Address {
	return &Address{
		schema:  schema,
		address: address,
	}
}

func ParseAddress(addressString string) (*Address, error) {
	parts := strings.Split(addressString, "://")

	if len(parts) != 2 {
		return nil, ErrInvalidAddressFormat
	}

	var schema AddressSchema

	switch strings.ToLower(parts[0]) {
	case "sms":
		schema = SMS
	case "email":
		schema = EMail
	default:
		return nil, ErrInvalidAddressFormat
	}

	return NewAddress(schema, parts[1]), nil
}

func (address Address) GetSchema() AddressSchema {
	return address.schema
}

func (address Address) GetAddress() string {
	return address.address
}

func (address Address) String() string {
	return string(address.schema) + "://" + address.address
}
