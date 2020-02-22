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
package mqtt

import (
	"context"
	"fmt"

	"github.com/cloustone/pandas/mixer/adaptors"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

type AdaptorFactory struct{}

func (r AdaptorFactory) Create(servingOptions *adaptors.AdaptorOptions) (adaptors.Adaptor, error) {
	return newMqttAdaptor(servingOptions)
}

type mqttAdaptor struct {
	context        context.Context
	shutdownFn     context.CancelFunc
	childRoutines  *errgroup.Group
	adaptorOptions *adaptors.AdaptorOptions
	mqttClient     mqtt.Client
}

func newMqttAdaptor(adaptorOptions *adaptors.AdaptorOptions) (adaptors.Adaptor, error) {
	broker := fmt.Sprintf("tcp://%s", adaptorOptions.ConnectURL)
	opts := mqtt.NewClientOptions().AddBroker(broker)
	opts.SetClientID(adaptors.BuildAdaptorID(adaptorOptions.Domain, adaptorOptions.Protocol))
	opts.SetUsername(adaptorOptions.Username)
	opts.SetPassword(adaptorOptions.Password)
	c := mqtt.NewClient(opts)

	if token := c.Connect(); token.Wait() && token.Error() != nil {
		logrus.WithError(token.Error())
		return nil, token.Error()
	}

	rootCtx, shutdownFn := context.WithCancel(context.Background())
	childRoutines, childCtx := errgroup.WithContext(rootCtx)

	s := &mqttAdaptor{
		context:        childCtx,
		shutdownFn:     shutdownFn,
		childRoutines:  childRoutines,
		adaptorOptions: adaptorOptions,
		mqttClient:     c,
	}

	return s, nil
}

func (r *mqttAdaptor) Options() *adaptors.AdaptorOptions { return r.adaptorOptions }

func (r *mqttAdaptor) handleReceivedMessage(client mqtt.Client, message mqtt.Message) {
	// payload := message.Payload()
}

func (r *mqttAdaptor) Start() error {
	for _, endpoint := range r.adaptorOptions.Endpoints {
		if token := r.mqttClient.Subscribe(endpoint, 0, r.handleReceivedMessage); token.Wait() && token.Error() != nil {
			logrus.WithError(token.Error())
			return token.Error()
		}
	}
	return nil
}

func (r *mqttAdaptor) GracefulShutdown() error {
	for _, endpoint := range r.adaptorOptions.Endpoints {
		if token := r.mqttClient.Unsubscribe(endpoint); token.Wait() && token.Error() != nil {
			logrus.WithError(token.Error())
			return token.Error()
		}
	}
	r.mqttClient.Disconnect(250)
	return nil
}

func (r *mqttAdaptor) WithMessageBuilder(adaptors.MessageBuilder) {}
