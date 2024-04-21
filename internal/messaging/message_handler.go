package messaging

type MessageHandler interface {
	Handle(Message) error
}
