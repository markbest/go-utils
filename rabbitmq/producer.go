package utils

import (
	"fmt"
	"github.com/streadway/amqp"
)

type Producer struct {
	amqpUrl      string
	exchange     string
	exchangeType string
	routingKey   string
}

// new producer
func NewProducer(amqpUrl, exchange, exchangeType, routingKey string) Producer {
	producer := Producer{
		amqpUrl:      amqpUrl,
		exchange:     exchange,
		exchangeType: exchangeType,
		routingKey:   routingKey,
	}
	return producer
}

// publish msg
func (p *Producer) Publish(body string) error {
	connection, err := amqp.Dial(p.amqpUrl)
	if err != nil {
		return fmt.Errorf("Dial: %s", err)
	}
	defer connection.Close()

	channel, err := connection.Channel()
	if err != nil {
		return fmt.Errorf("Channel: %s", err)
	}
	defer channel.Close()

	if err := channel.ExchangeDeclare(
		p.exchange,     // name
		p.exchangeType, // type
		true,           // durable
		false,          // auto-deleted
		false,          // internal
		false,          // noWait
		nil,            // arguments
	); err != nil {
		return fmt.Errorf("Exchange Declare: %s", err)
	}

	if err := channel.Publish(
		p.exchange,   // publish to an exchange
		p.routingKey, // routing to 0 or more queues
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			Headers:         amqp.Table{},
			ContentType:     "text/plain",
			ContentEncoding: "",
			Body:            []byte(body),
			DeliveryMode:    amqp.Transient, // 1=non-persistent, 2=persistent
			Priority:        0,              // 0-9
		},
	); err != nil {
		return fmt.Errorf("Exchange Publish: %s", err)
	}
	return nil
}
