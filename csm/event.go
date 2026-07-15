package csm

import "fmt"

type EventChannel int

const (
	Internal EventChannel = iota
	External
	Peripheral
)

type Event struct {
	Topic   string
	Channel EventChannel
	Data    []ContextVariable
	Target  string
	Source  string
}

type ConditionalEvent struct {
	Event
	Provided Expression
}

func (c EventChannel) String() string {
	switch c {
	case Internal:
		return "Internal"
	case External:
		return "External"
	case Peripheral:
		return "Peripheral"
	default:
		return fmt.Sprintf("UnknownChannel(%d)", c)
	}
}

func (e Event) String() string {
	return fmt.Sprintf("Event(topic=%q, channel=%s, target=%q, source=%q)",
		e.Topic, e.Channel, e.Target, e.Source)
}

func (e ConditionalEvent) String() string {
	return fmt.Sprintf("ConditionalEvent(provided=%q, event=%+v)", e.Provided, e.Event)
}
