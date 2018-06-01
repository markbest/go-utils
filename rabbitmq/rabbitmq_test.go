package utils

import (
	"testing"
)

var (
	amqpUrl      = "amqp://guest:guest@127.0.0.1:5672/"
	exchange     = "test_exchange"
	exchangeType = "direct"
	routingKey   = "test"
	queueName    = "test"
)

func TestRabbitmq(t *testing.T) {
	producer := NewProducer(amqpUrl, exchange, exchangeType, routingKey)
	err := producer.Publish("This is a test")
	if err != nil {
		t.Fatalf(err.Error())
	}

	flag := make(chan bool)
	c, err := NewConsumer(amqpUrl, exchange, exchangeType, queueName, routingKey, "")
	if err != nil {
		t.Fatalf(err.Error())
	}
	c.Receive()
	<-flag
}
