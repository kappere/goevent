package goevent

import (
	"container/list"
	"sync"
)

var queue = newEventQueue()

type EventQueue struct {
	*list.List
	lock    sync.Mutex
	c_empty chan int
}

func newEventQueue() *EventQueue {
	return &EventQueue{
		List:    list.New().Init(),
		lock:    sync.Mutex{},
		c_empty: make(chan int, 1),
	}
}

func (q *EventQueue) Offer(msg EventInterface) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.PushBack(msg)
	if len(q.c_empty) == 0 {
		q.c_empty <- 1
	}
}

func (q *EventQueue) Take() EventInterface {
	q.lock.Lock()
	defer q.lock.Unlock()
	if len(q.c_empty) == 1 {
		<-q.c_empty
	}
	if q.Len() == 0 {
		q.lock.Unlock()
		<-q.c_empty
		q.lock.Lock()
	}
	return q.Remove(q.Front()).(EventInterface)
}

func offerEvent(event EventInterface) {
	queue.Offer(event)
}

func takeEvent() EventInterface {
	return queue.Take()
}
