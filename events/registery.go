package events

import (
	"kook-bot-chatgpt/config"
	"kook-bot-chatgpt/events/handler"
)

// 创建一个事件调度器并且设置一个指令阻塞队列100长度，和并发10
var Dispatcher = NewEventDispatcher(config.GlobalConfig.Queue.MaxLength, config.GlobalConfig.Concurrency.MaxGoroutines)

// 注册
func Init() {

	// 注册事件处理handler
	Dispatcher.AddHandler("/help", &handler.Help{})
	Dispatcher.AddHandler("/chat", &handler.Chat{})
	Dispatcher.AddHandler("/reset", &handler.Reset{})

	// 启动指令调度器
	go Dispatcher.Run()
}
