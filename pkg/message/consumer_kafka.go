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

import (
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/Shopify/sarama"
	"github.com/golang/glog"
)

type kafkaConsumer struct {
	khosts      string                 // kafka server list
	subscribers map[string]*subscriber // kafka client endpoint
	mutex       sync.Mutex
	clientId    string
	msgFactory  MessageFactory
}

type subscriber struct {
	topic      string
	waitgroup  sync.WaitGroup
	handler    MessageHandlerFunc
	ctx        interface{}
	consumer   sarama.Consumer
	clientId   string
	pconsumers []sarama.PartitionConsumer
	quitChan   chan interface{}
}

func newKafkaConsumer(c *Config) (Consumer, error) {
	khosts := c.Hosts
	if khosts == "" {
		return nil, errors.New("message service is not rightly configed")
	}
	names := strings.Split(khosts, ":")
	if len(names) == 1 {
		khosts = khosts + ":9092"
	}
	return &kafkaConsumer{
		khosts:      khosts,
		subscribers: make(map[string]*subscriber),
		mutex:       sync.Mutex{},
	}, nil
}

func (p *kafkaConsumer) WithClientId(clientId string)             { p.clientId = clientId }
func (p *kafkaConsumer) SetMessageFactory(factory MessageFactory) { p.msgFactory = factory }

func (p *kafkaConsumer) Subscribe(topic string, queueName string, handler MessageHandlerFunc, ctx interface{}) error {
	if _, found := p.subscribers[topic]; found {
		return fmt.Errorf("topic '%s' already subcribed", topic)
	}
	clientId := fmt.Sprintf("%s_%d", p.clientId, len(p.subscribers))
	config := sarama.NewConfig()
	config.ClientID = clientId
	// config.kafkaConsumer.MaxWaitTime = time.Duration(5 * time.Second)
	// config.kafkaConsumer.Offsets.CommitInterval = 1 * time.Second
	// config.kafkaConsumer.Offsets.Initial = sarama.OffsetNewest
	consumer, err := sarama.NewConsumer(strings.Split(p.khosts, ","), config)
	if err != nil {
		return fmt.Errorf("host = '%s', %s", p.khosts, err.Error())
	}
	p.subscribers[topic] = &subscriber{
		topic:      topic,
		waitgroup:  sync.WaitGroup{},
		handler:    handler,
		ctx:        ctx,
		consumer:   consumer,
		clientId:   clientId,
		pconsumers: []sarama.PartitionConsumer{},
		quitChan:   make(chan interface{}, 1),
	}
	return nil
}

func (p *kafkaConsumer) UnSubscribe(topic string) error {
	return nil
}

func (p *kafkaConsumer) Start() error {
	for topic, sub := range p.subscribers {
		consumer := sub.consumer
		if partitionList, err := consumer.Partitions(topic); err == nil {
			for partition := range partitionList {
				pc, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
				if err != nil {
					return err
				}
				sub.pconsumers = append(sub.pconsumers, pc)
				sub.waitgroup.Add(1)
				go func(s *subscriber, pc sarama.PartitionConsumer) {
					defer s.waitgroup.Done()
					for {
						select {
						case <-s.quitChan:
							return
						case err := <-pc.Errors():
							glog.Error(err)
							//return
						case msg := <-pc.Messages():
							if s.handler != nil {
								t := p.msgFactory.CreateMessage(topic)
								if t != nil && t.Deserialize(msg.Value, JSONSerialization) == nil {
									s.handler(t, s.ctx)
								} else {
									glog.Errorf("topic '%s' handle failure", topic)
								}
							}
						}
					}
				}(sub, pc)
			}
		}
	}
	return nil
}

func (p *kafkaConsumer) Close() {
	for _, sub := range p.subscribers {
		sub.quitChan <- true
		for _, pc := range sub.pconsumers {
			pc.Close()
		}
		sub.consumer.Close()
		sub.waitgroup.Wait()
	}
}
