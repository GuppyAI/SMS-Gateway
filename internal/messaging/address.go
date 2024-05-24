package messaging

type AddressSchema string

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

func (address Address) GetSchema() AddressSchema {
	return address.schema
}

func (address Address) GetAddress() string {
	return address.address
}

func (address Address) String() string {
	return string(address.schema) + "://" + address.address
}
