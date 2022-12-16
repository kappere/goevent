package goevent

import (
	"testing"
	"time"
)

type MyEvent struct {
	Event
}

func TestPublish(t *testing.T) {
	Async(true)
	Sequential(true)
	Subscribe(MyEvent{}, func(event EventInterface) {
		t.Logf("MyEvent: %s", event.String())
	})
	Subscribe(Event{}, func(event EventInterface) {
		t.Logf("Event: %s", event.String())
	})
	Subscribe(Event{}, func(event EventInterface) {
		t.Logf("Event2: %s", event.String())
	})
	Publish(MyEvent{
		Event: NewEvent("0"),
	})
	Publish(NewEvent("1"))
	Publish(NewEvent("2"))
	Publish(NewEvent("3"))
	Publish(NewEvent("4"))
	Publish(NewEvent("5"))
	Publish(NewEvent("6"))
	Publish(NewEvent("7"))
	Publish(NewEvent("8"))
	time.Sleep(1 * time.Second)
}
