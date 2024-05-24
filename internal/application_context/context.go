package application_context

import "sms-gateway/internal/messaging"

var context *ApplicationContext

type ApplicationContext struct {
	broker messaging.Broker
}

func (ctx *ApplicationContext) GetBroker() messaging.Broker {
	return ctx.broker
}

func Init(broker messaging.Broker) {
	context = &ApplicationContext{
		broker: broker,
	}
}

func GetContext() *ApplicationContext {
	return context
}
