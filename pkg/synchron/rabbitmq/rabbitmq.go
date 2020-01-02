//  Licensed under the Apache License, Version 2.0 (the "License"); you may
//  not use p file except in compliance with the License. You may obtain
//  a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
//  WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
//  License for the specific language governing permissions and limitations
//  under the License.
package rabbitmq

import (
	"encoding/json"
	"fmt"

	"github.com/cloustone/pandas/pkg/synchron"

	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

const NAME = "rabbitmq"

type RabbitmqSynchronizer struct {
	observers   []synchron.Observer
	conn        *amqp.Connection
	ch          *amqp.Channel
	subscribers map[string]subscriber
}

type subscriber struct {
	topic    string
	clientId string
	queue    amqp.Queue
	observer synchron.Observer
}

func NewSynchronizer(usr string, pwd string, hosts string) synchron.Synchronizer {
	connectUrl := fmt.Sprintf("amqp://%s:%s@%s/", usr, pwd, hosts)
	conn, err := amqp.Dial(connectUrl)
	if err != nil {
		logrus.Fatalf(err.Error())
	}
	ch, err := conn.Channel()
	if err != nil {
		logrus.Fatalf(err.Error())
	}
	return &RabbitmqSynchronizer{
		observers:   []synchron.Observer{},
		conn:        conn,
		ch:          ch,
		subscribers: make(map[string]subscriber),
	}
}

func (r *RabbitmqSynchronizer) AsMember() {}

func (r *RabbitmqSynchronizer) WithRootPath(path string) synchron.Synchronizer { return r }
func (r *RabbitmqSynchronizer) Notify(n synchron.Notification) {
	body, err := json.Marshal(&n)
	if err != nil {
		logrus.WithError(err)
	}
	pub := amqp.Publishing{
		ContentType: "text/json",
		Body:        body,
	}
	if err := r.ch.Publish(n.Path, "", false, false, pub); err != nil {
		logrus.WithError(err)
	}
}

func (r *RabbitmqSynchronizer) RegisterObserver(path string, obs synchron.Observer) {
	if err := r.ch.ExchangeDeclare(path, "fanout", true, false, false, false, nil); err != nil {
		logrus.WithError(err).Fatalf("register observer '%s' failed", path)
	}
	queue, err := r.ch.QueueDeclare(
		path,  // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,
	)
	if err != nil {
		logrus.WithError(err).Fatalf("register observer '%s' failed", path)
	}
	if err := r.ch.QueueBind(path, "", path, false, nil); err != nil {
		logrus.WithError(err).Fatalf("register observer '%s' failed", path)
	}
	sub := subscriber{
		topic:    path,
		clientId: path,
		queue:    queue,
		observer: obs,
	}
	msgs, err := r.ch.Consume(
		path,  // queue
		"",    // consumer
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		logrus.WithError(err).Fatalf("register observer '%s' failed", path)
	}
	go func(r *RabbitmqSynchronizer, sub subscriber) {
		for msg := range msgs {
			n := synchron.Notification{}
			if err := json.Unmarshal(msg.Body, &n); err != nil {
				logrus.WithError(err)
			}
			sub.observer.OnSynchronizationNotified(r, n)
		}
	}(r, sub)
	r.subscribers[path] = sub
}
