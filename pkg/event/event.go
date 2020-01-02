package event

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/cloustone/pandas/pkg/message"
)

const (
	FmtOfMqttEventBus   = "iot-%s-event-mqtt"
	FmtOfBrokerEventBus = "iot-%s-event-broker"
)

const (
	SessionCreate    = 1
	SessionDestroy   = 2
	TopicPublish     = 3
	TopicSubscribe   = 4
	TopicUnsubscribe = 5
	QutoChange       = 6
	SessionResume    = 7
	AuthChange       = 8
)

type Event interface {
	SetBrokerId(brokerId string)
	SetType(eventType uint32)
	SetClientId(clientID string)
	SetDeviceName(deviceName string)
	SetProductId(productId string)
	GetBrokerId() string
	GetType() uint32
	GetClientId() string
	GetDeviceName() string
	GetProductId() string
}

// NameOfEvent return event name
func NameOfEvent(t uint32) string {
	switch t {
	case SessionCreate:
		return "SessionCreate"
	case SessionDestroy:
		return "SessionDestroy"
	case TopicPublish:
		return "TopicPublish"
	case TopicSubscribe:
		return "TopicSubscribe"
	case TopicUnsubscribe:
		return "TopicUnsubscribe"
	case QutoChange:
		return "QutoChange"
	case SessionResume:
		return "SessionResume"
	case AuthChange:
		return "AuthChange"
	default:
		return "Unknown"
	}
}

// FullNameOfEvent return event information
func FullNameOfEvent(e Event) string {
	return fmt.Sprintf("Event:%s, broker:%s, clientid:%s", NameOfEvent(e.GetType()), e.GetBrokerId(), e.GetClientId())
}

type CodecOption struct {
	Format string
}

var (
	JSONCodec = CodecOption{Format: "json"}
)

// Decode unmarshal event from raw buffer using rawEvent
func Decode(msg message.Message, opt CodecOption) (Event, error) {
	re, ok := msg.(*message.Broker)
	if !ok || re == nil {
		return nil, errors.New("invalid broker event")
	}
	switch re.EventType {
	case SessionCreate:
		e := &SessionCreateEvent{}
		err := json.Unmarshal(re.Payload, e)
		return e, err
	case SessionDestroy:
		e := &SessionDestroyEvent{}
		err := json.Unmarshal(re.Payload, e)
		return e, err
	case SessionResume:
		e := &SessionResumeEvent{}
		err := json.Unmarshal(re.Payload, e)
		return e, err
	case TopicSubscribe:
		e := &TopicSubscribeEvent{}
		err := json.Unmarshal(re.Payload, e)
		return e, err
	case TopicUnsubscribe:
		e := &TopicUnsubscribeEvent{}
		err := json.Unmarshal(re.Payload, e)
		return e, err
	case TopicPublish:
		e := &TopicPublishEvent{}
		err := json.Unmarshal(re.Payload, e)
		return e, err
	case QutoChange:
		e := &QutoChangeEvent{}
		err := json.Unmarshal(re.Payload, e)
		return e, err
	case AuthChange:
		e := &AuthChangeEvent{}
		err := json.Unmarshal(re.Payload, e)
		return e, err
	default:
		return nil, fmt.Errorf("invalid event type '%d'", re.EventType)
	}
}

// Encode serialize event using rawEvent
func Encode(e Event, opt CodecOption) ([]byte, error) {
	switch opt {
	case JSONCodec:
		return json.Marshal(e)
	}

	return nil, fmt.Errorf("invalid codec option")
}
