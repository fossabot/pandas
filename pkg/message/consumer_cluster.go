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
	"time"

	"github.com/Shopify/sarama"
	"github.com/bsm/sarama-cluster"
	"github.com/golang/glog"
)

// consumer
var consumers map[string]*cluster.Consumer

func StartConsumerCluster(khosts string, productKey string, topics []string) error {
	if consumers == nil {
		consumers = make(map[string]*cluster.Consumer)
	}
	if consumer, ok := consumers[productKey]; ok {
		consumer.Close()
	}
	if khosts == "" {
		return errors.New("message service is not rightly configed")
	}
	groupID := productKey
	config := cluster.NewConfig()
	config.Group.Return.Notifications = true
	config.Consumer.Offsets.CommitInterval = 1 * time.Second
	config.Consumer.Offsets.Initial = sarama.OffsetNewest //

	c, err := cluster.NewConsumer(strings.Split(khosts, ","), groupID, topics, config)
	if err != nil {
		glog.Errorf("Failed open consumer: %v", err)
		return err
	}
	consumers[productKey] = c
	defer c.Close()
	go func(c *cluster.Consumer) {
		errors := c.Errors()
		noti := c.Notifications()
		for {
			select {
			case err := <-errors:
				glog.Errorln(err)
			case <-noti:
			}
		}
	}(c)

	for msg := range c.Messages() {
		////////////////////////////////////////////////////////////////////////
		// add device topic process.
		// Subscribe Topic /{productKey}/{deviceId}/update
		// Public    Topic /{productKey}/{deviceId}/update
		fmt.Printf("hello:%s/%d/%d\t%s\n", msg.Topic, msg.Partition, msg.Offset, msg.Value)
		c.MarkOffset(msg, "") //MarkOffset
	}
	return nil
}

func CloseConsumerCluster() {
	for _, c := range consumers {
		c.Close()
	}

	glog.Info("Device engine end!")
}
