package messaging

type Kind int

const (
	Request  Kind = 0
	Response      = 1
)

func (kind Kind) String() string {
	switch kind {
	case 0:
		return "Request"
	case 1:
		return "Response"
	default:
		return "Unknown"
	}
}

// Message is used to model messages
type Message struct {
	kind    Kind
	address Address
	body    string
}

// NewMessage constructs a new message
func NewMessage(kind Kind, address Address, message string) *Message {
	return &Message{kind, address, message}
}

// GetAddress gives back the sender address of the message
func (message *Message) GetAddress() Address {
	return message.address
}

// GetMessageBody returns the body of the message
func (message *Message) GetMessageBody() string {
	return message.body
}
