package main

import (
	"kook-bot-chatgpt/config"
	"kook-bot-chatgpt/controllers"
	"kook-bot-chatgpt/events"
	"kook-bot-chatgpt/middlewares"
	"kook-bot-chatgpt/utils"
	"fmt"
	"log"
	"net/http"
)

func main() {

	// 创建中间件链
	middleware := []func(http.Handler) http.Handler{
		middlewares.LoggingMiddleware,
		middlewares.ValidateMiddleware,
		middlewares.DecompressMiddleware,
	}

	// webkook请求
	webhook := http.HandlerFunc(controllers.WebhookHandler)

	// 注册 HTTP 服务器处理函数
	http.Handle("/", utils.ChainMiddleware(webhook, middleware...))

	// 注册事件
	events.Init()
	
	// 启动 HTTP 服务器
	log.Println("Server started")
	s := fmt.Sprintf("%v:%v", config.GlobalConfig.Server.Host, config.GlobalConfig.Server.Port)
	log.Fatal(http.ListenAndServe(s, nil))
}
