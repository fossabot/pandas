package message

import (
	"fmt"
	"sync"

	logr "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type rabbitConsumer struct {
	msgFactory  MessageFactory
	conn        *amqp.Connection
	ch          *amqp.Channel
	khosts      string                  // kafka server list
	subscribers map[string]*rsubscriber // kafka client endpoint
	mutex       sync.Mutex
	clientID    string
}

type rsubscriber struct {
	topic     string
	waitgroup sync.WaitGroup
	handler   MessageHandlerFunc
	ctx       interface{}
	clientID  string
	quitChan  chan interface{}
	queue     amqp.Queue
}

func newRabbitConsumer(c *Config) (Consumer, error) {
	url := fmt.Sprintf("amqp://%s:%s@%s/", c.User, c.Password, c.Hosts)
	logr.Debugf("create rabbit consumer, rabbitmq server url : %s", url)

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
	return &rabbitConsumer{
		conn:        conn,
		ch:          ch,
		subscribers: make(map[string]*rsubscriber),
		mutex:       sync.Mutex{},
	}, nil
}

func (r *rabbitConsumer) SetMessageFactory(factory MessageFactory) { r.msgFactory = factory }
func (r *rabbitConsumer) WithClientId(clientID string)             { r.clientID = clientID }

func (r *rabbitConsumer) Subscribe(topic string, queueName string, handler MessageHandlerFunc, ctx interface{}) error {
	logr.Debugf("rabbitmq consumer subscribe a queue with topic '%s'", topic)

	if _, found := r.subscribers[topic]; found {
		return fmt.Errorf("topic '%s' already subcribed", topic)
	}

	clientID := fmt.Sprintf("%s_%d", r.clientID, len(r.subscribers))
	var q amqp.Queue
	var err error

	if err = r.ch.ExchangeDeclare(topic, "fanout", true, false, false, false, nil); err != nil {
		logr.WithError(err).Errorf("declare exchange with topic '%s' failed", topic)
		return err
	}
	q, err = r.ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when usused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		logr.WithError(err).Errorf("create queue '%s' failed", topic)
		return err
	}
	err = r.ch.QueueBind(
		q.Name, // queue name
		"",     // routing key
		topic,  // exchange
		false,
		nil)

	r.subscribers[topic] = &rsubscriber{
		topic:     topic,
		waitgroup: sync.WaitGroup{},
		handler:   handler,
		ctx:       ctx,
		clientID:  clientID,
		quitChan:  make(chan interface{}, 1),
		queue:     q,
	}

	return nil
}

func (r *rabbitConsumer) UnSubscribe(topic string) error {
	return r.ch.ExchangeDelete(
		topic,
		false,
		true,
	)
}

func (r *rabbitConsumer) Start() error {
	logr.Debugf("start rabbit consumer '%s'", r.clientID)

	for _, sub := range r.subscribers {
		// create consumer
		msgs, err := r.ch.Consume(
			sub.queue.Name, // queue
			"",             // consumer
			true,           // auto-ack
			false,          // exclusive
			false,          // no-local
			false,          // no-wait
			nil,            // args
		)
		if err != nil {
			r.ch.QueueDelete(sub.queue.Name, true, true, true)
			logr.WithError(err).Errorf("create consumer for '%s' failed", sub.topic)
			return err
		}
		go func(r *rabbitConsumer, sub *rsubscriber) {
			for d := range msgs {
				t := r.msgFactory.CreateMessage(sub.topic)
				if t != nil && t.Deserialize(d.Body, JSONSerialization) == nil {
					sub.handler(t, sub.ctx)
				} else {
					logr.Errorf("rabbitmq consumer failed to receive message")
				}
			}
		}(r, sub)
	}

	return nil
}

func (r *rabbitConsumer) Close() {
	r.conn.Close()
}
