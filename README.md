# goevent
简单golang事件处理工具
### 快速开始
```go
package main

import (
	"fmt"
	"sync"

	"github.com/kappere/goevent"
)

type MyEvent struct {
	goevent.Event
	Name string
}

func main() {
	wg := sync.WaitGroup{}

	// 事件是否异步执行（默认异步）
	goevent.Async(true)

	// 是否顺序消费事件（默认否）
	goevent.Sequential(false)

	// 日志记录函数（默认fmt.Printf打印）
	goevent.LoggerFunc(func(s string) {
		fmt.Println(s)
	})

	// 订阅事件
	goevent.Subscribe(MyEvent{}, func(event goevent.EventInterface) {
		fmt.Printf("Consume event: %s\n", event.(MyEvent).Name)
		wg.Done()
	})

	wg.Add(1)
	// 发布事件
	goevent.Publish(MyEvent{
		Event: goevent.NewEvent("main"),
		Name:  "My first event",
	})

	// 等待事件消费结束
	wg.Wait()
}

```
