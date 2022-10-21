package mqtt

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
)

// Deprecated: Use mqtt.Consumer instead.
type Consumer struct {
	amqpCon *amqp.Connection
	ch      *amqp.Channel
	prefix  string
	group   string
}

// Deprecated: Use mqtt.NewConsumer instead.
func NewConsumer(amqpCon *amqp.Connection, prefix string, group string) *Consumer {
	return &Consumer{amqpCon: amqpCon, prefix: prefix, group: group}
}

func (b *Consumer) getChannel() (*amqp.Channel, error) {
	if b.ch == nil {
		channel, err := b.amqpCon.Channel()
		if err != nil {
			return nil, err
		}
		b.ch = channel
	} else if b.ch.IsClosed() {
		channel, err := b.amqpCon.Channel()
		if err != nil {
			return nil, err
		}
		b.ch = channel
	}

	return b.ch, nil
}

func (b *Consumer) BroadcastMessage(log *logrus.Entry, channelName string, onMessage func(data amqp.Delivery) error) error {
	channel, err := b.getChannel()
	if err != nil {
		return err
	}
	if err := channel.ExchangeDeclare(
		fmt.Sprintf("%s%s", b.prefix, channelName), // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	); err != nil {
		return err
	}

	q, err := channel.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return err
	}

	if err := channel.QueueBind(
		q.Name, // queue name
		"",     // routing key
		fmt.Sprintf("%s%s", b.prefix, channelName), // exchange
		false,
		nil,
	); err != nil {
		return err
	}

	msgs, err := channel.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			if err := onMessage(d); err != nil {
				log.Error(err)
				d.Nack(false, true)
			} else {
				d.Ack(false)
			}
		}
	}()

	return nil
}

func (b *Consumer) QueueMessage(log *logrus.Entry, channelName string, onMessage func(data amqp.Delivery) error) error {
	channel, err := b.getChannel()
	if err != nil {
		return err
	}
	if err := channel.ExchangeDeclare(
		fmt.Sprintf("%s%s", b.prefix, channelName), // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	); err != nil {
		return err
	}

	q, err := channel.QueueDeclare(
		fmt.Sprintf("%s%s.queue.%s", b.prefix, channelName, b.group), // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return err
	}

	if err := channel.QueueBind(
		q.Name, // queue name
		"",     // routing key
		fmt.Sprintf("%s%s", b.prefix, channelName), // exchange
		false,
		nil,
	); err != nil {
		return err
	}

	msgs, err := channel.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			if err := onMessage(d); err != nil {
				log.Error(err)
				d.Nack(false, true)
			} else {
				d.Ack(false)
			}
		}
	}()

	return nil
}
