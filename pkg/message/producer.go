//  Licensed under the Apache License, Version 2.0 (the "License"); you may
//  not use this file except in compliance with the License. You may obtain
//  a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
//  WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
//  License for the specific language governing permissions and limitations
//  under the License.
package message

import "fmt"

type Producer interface {
	WithClientId(clientId string)
	SendMessage(msg Message) error
	SendMessages(msgs []Message) error
	Close()
}

func NewProducer(c *Config, sync bool) (Producer, error) {
	switch c.Backend {
	case MessageBackendKafka:
		return newKafkaProducer(c, sync)

	case MessageBackendRabbit:
		return newRabbitProducer(c, sync)
	}
	return nil, fmt.Errorf("invalid message backend '%s'", c.Backend)
}

func PostMessage(msg Message, c *Config) (err error) {
	if producer, err := NewProducer(c, false); err == nil {
		defer producer.Close()
		return producer.SendMessage(msg)
	}
	return
}

func PostMessages(msgs []Message, c *Config) (err error) {
	if producer, err := NewProducer(c, false); err == nil {
		defer producer.Close()
		return producer.SendMessages(msgs)
	}
	return
}

func SendMessage(msg Message, c *Config) (err error) {
	if producer, err := NewProducer(c, true); err == nil {
		defer producer.Close()
		return producer.SendMessage(msg)
	}
	return
}

func SendMessages(msgs []Message, c *Config) (err error) {
	if producer, err := NewProducer(c, true); err == nil {
		defer producer.Close()
		return producer.SendMessages(msgs)
	}
	return
}
