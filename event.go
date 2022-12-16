package goevent

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type EventInterface interface {
	Id() string
	Timestamp() time.Time
	Source() interface{}
	String() string
}

type Event struct {
	id        string
	timestamp time.Time
	source    interface{}
}

func (e Event) Id() string {
	return e.id
}

func (e Event) Timestamp() time.Time {
	return e.timestamp
}

func (e Event) Source() interface{} {
	return e.source
}

func (e Event) String() string {
	return fmt.Sprintf("Event(id=%s,timestamp=%d,source=%v)", e.id, e.timestamp.Unix(), e.source)
}

func NewEvent(source interface{}) Event {
	return Event{
		id:        uuid.NewString(),
		timestamp: time.Now(),
		source:    source,
	}
}
