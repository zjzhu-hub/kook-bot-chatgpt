package events

import (
	"kook-bot-chatgpt/queues"
)

//定义事件分发器
type EventDispatcher struct {
	queue           *queues.ConcurrentQueue 	// 队列实例
    maxConcurrent   int                         // 最大并发数
    semaphore       chan struct{}               // 控制并发的信号量
	handlerMap     	map[string]EventHandler 	// key 指令名称 value 指令需要执行的handler 
}

//定义事件处理程序接口
type EventHandler interface {
    Handler(body map[string]interface{}) 
}

func (ed *EventDispatcher) AddHandler(command string, handler EventHandler) {
    ed.handlerMap[command] = handler
}

func (ed *EventDispatcher) Handle(command string, body map[string]interface{}) {
    handler, ok := ed.handlerMap[command]
    if !ok {
        return
    }
    handler.Handler(body)
}

func (ed *EventDispatcher) Push(c queues.RequestWithCommand){
    ed.queue.Push(c)
}

func (ed *EventDispatcher) Run() {
    for {
        // 往channel 中发送一个空的 struct 类型的值， 这行代码会阻塞直到 channel 中有空闲的空间；相当于占个坑 直到有下一个坑释放才能写入
        ed.semaphore <- struct{}{}

        // Pop 出队列中的第一个元素
        c := ed.queue.Pop()

        // 启动一个新的 goroutine 执行任务
        go func() {
            ed.Handle(c.Command, c.Body)

            // 释放一个信号量
            <-ed.semaphore
        }()
    }
}

func NewEventDispatcher(queueCapacity, maxConcurrent int) *EventDispatcher {
    return &EventDispatcher{
        handlerMap: 	make(map[string]EventHandler),
		queue:         	queues.NewConcurrentQueue(queueCapacity),
        maxConcurrent:  maxConcurrent,
        semaphore:      make(chan struct{}, maxConcurrent),
    }
}
