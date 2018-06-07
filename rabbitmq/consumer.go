package utils

import (
	"fmt"
	"github.com/streadway/amqp"
)

type Consumer struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   amqp.Queue
	tag     string
	done    chan error
}

// new consumer
func NewConsumer(amqpUrl, exchange, exchangeType, queueName, key, tag string) (c *Consumer, err error) {
	c = &Consumer{
		conn:    nil,
		channel: nil,
		tag:     tag,
		done:    make(chan error),
	}

	c.conn, err = amqp.Dial(amqpUrl)
	if err != nil {
		return nil, fmt.Errorf("dial: %s", err)
	}

	c.channel, err = c.conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("channel: %s", err)
	}

	if err = c.channel.ExchangeDeclare(
		exchange,     // name of the exchange
		exchangeType, // type
		true,         // durable
		false,        // delete when complete
		false,        // internal
		false,        // noWait
		nil,          // arguments
	); err != nil {
		return nil, fmt.Errorf("exchange Declare: %s", err)
	}

	c.queue, err = c.channel.QueueDeclare(
		queueName, // name of the queue
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // noWait
		nil,       // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("queue Declare: %s", err)
	}

	if err = c.channel.QueueBind(
		c.queue.Name, // name of the queue
		key,          // bindingKey
		exchange,     // sourceExchange
		false,        // noWait
		nil,          // arguments
	); err != nil {
		return nil, fmt.Errorf("queue Bind: %s", err)
	}
	return c, nil
}

// shut down consumer
func (c *Consumer) Shutdown() error {
	if err := c.channel.Cancel(c.tag, true); err != nil {
		return fmt.Errorf("consumer cancel failed: %s", err)
	}

	if err := c.conn.Close(); err != nil {
		return fmt.Errorf("AMQP connection close error: %s", err)
	}
	return <-c.done
}

// monitor get queue msg
func (c *Consumer) Receive() {
	deliveries, err := c.channel.Consume(
		c.queue.Name, // name
		c.tag,        // consumerTag,
		false,        // noAck
		false,        // exclusive
		false,        // noLocal
		false,        // noWait
		nil,          // arguments
	)
	if err != nil {
		panic(err)
	}

	go func() {
		for d := range deliveries {
			fmt.Println("got msg: ", string(d.Body))
			d.Ack(false)
		}
		c.done <- nil
	}()
}
