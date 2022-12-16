package goevent

import (
	"fmt"
	"reflect"
	"sync"
)

type ConsumerApplyer = func(event EventInterface)

var (
	consumers                     = make(map[reflect.Type][]ConsumerApplyer)
	logFunc          func(string) = func(s string) { fmt.Println(s) }
	consumer_lock                 = sync.Mutex{}
	consumer_running bool         = false
	sequential                    = false
)

// 日志记录函数（默认fmt.Printf打印）
func LoggerFunc(f func(string)) {
	logFunc = f
}

// 是否顺序消费事件（默认否）
func Sequential(seq bool) {
	sequential = seq
}

// 订阅事件
func Subscribe(event_type_obj EventInterface, apply ConsumerApplyer) {
	subscribeWithQueue(event_type_obj, apply, consumers)
}

func subscribeWithQueue(event_type_obj EventInterface, apply ConsumerApplyer, consumer_queue map[reflect.Type][]ConsumerApplyer) {
	consumer_lock.Lock()
	defer consumer_lock.Unlock()
	event_type := reflect.TypeOf(event_type_obj)
	applys, exists := consumer_queue[event_type]
	if !exists {
		applys = make([]ConsumerApplyer, 0)
	}
	consumer_queue[event_type] = append(applys, apply)
}

func findConsumers(event_type reflect.Type) []ConsumerApplyer {
	applys, exists := consumers[event_type]
	if !exists {
		return []ConsumerApplyer{}
	}
	return applys
}

func runConsumer() {
	if consumer_running {
		return
	}
	consumer_lock.Lock()
	defer consumer_lock.Unlock()
	if consumer_running {
		return
	}
	consumer_running = true
	go func() {
		for {
			event := takeEvent()
			consumeAll(event)
		}
	}()
}

func consumeAll(event EventInterface) {
	event_type := reflect.TypeOf(event)
	applys := findConsumers(event_type)
	for _, apply := range applys {
		if async && !sequential {
			go consume(apply, event)
		} else {
			consume(apply, event)
		}
	}
}

func consume(apply ConsumerApplyer, event EventInterface) {
	defer func() {
		if err := recover(); err != nil {
			logFunc(fmt.Sprintf("%v", err))
		}
	}()
	apply(event)
}
