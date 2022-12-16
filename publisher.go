package goevent

var (
	async = true
)

// 事件是否异步执行（默认异步）
func Async(as bool) {
	async = as
}

// 发布事件
func Publish(event EventInterface) {
	if async {
		runConsumer()
		offerEvent(event)
	} else {
		consumeAll(event)
	}
}
