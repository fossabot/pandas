package message

import (
	"fmt"

	logr "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type rabbitProducer struct {
	config   *Config
	sync     bool
	clientId string
	conn     *amqp.Connection
	ch       *amqp.Channel
}

func newRabbitProducer(c *Config, sync bool) (Producer, error) {
	url := fmt.Sprintf("amqp://%s:%s@%s/", c.User, c.Password, c.Hosts)
	logr.Debugf("create rabbit producer, rabbitmq server url : %s", url)

	conn, err := amqp.Dial(url)
	if err != nil {
		logr.WithError(err).Errorf("connect with rabbitmq server failed")
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		logr.WithError(err).Errorf("failed to create a rabbitmq channel")
		return nil, err
	}

	r := &rabbitProducer{
		config: c,
		sync:   sync,
		conn:   conn,
		ch:     ch,
	}
	return r, nil
}

func (r *rabbitProducer) WithClientId(clientId string) { r.clientId = clientId }

func (r *rabbitProducer) SendMessage(msg Message) error {
	topic := msg.Topic()
	body, err := msg.Serialize(JSONSerialization)
	if err != nil {
		logr.WithError(err).Errorf("serialize message '%s' failed", topic)
		return err
	}

	if err := r.ch.ExchangeDeclare(topic, "fanout", true, false, false, false, nil); err != nil {
		logr.WithError(err).Errorf("declare exchange with topic '%s' failed", topic)
		return err
	}

	pub := amqp.Publishing{
		ContentType: "text/json",
		Body:        []byte(body),
	}
	logr.Infoln("broadcast begin to publish event to rabbitmq, topic:", topic)
	if err := r.ch.Publish(topic, "", false, false, pub); err != nil {
		logr.WithError(err).Errorf("broadcast message on topic '%s' failed", topic)
	}
	return nil
}

func (r *rabbitProducer) SendMessages(msgs []Message) error {
	for _, msg := range msgs {
		if err := r.SendMessage(msg); err != nil {
			return err
		}
	}
	return nil
}

func (r *rabbitProducer) Close() {
	r.conn.Close()
}
